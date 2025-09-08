package services

import (
	"context"
	guidanceTemplateRepository "scs-operator/internal/app/guidance-template/repository"
	"scs-operator/internal/app/incident/dto"
	repo "scs-operator/internal/app/incident/repository"
	userRepositories "scs-operator/internal/app/user/repository"
	"scs-operator/internal/models"
	"scs-operator/internal/types"
	"scs-operator/pkg/errors"
	kafka_client "scs-operator/pkg/kafka"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type Service struct {
	incidentRepo             repo.IncidentRepository
	incidentGuidanceRepo     repo.IncidentGuidanceRepository
	incidentGuidanceStepRepo repo.IncidentGuidanceStepRepository
	userRepo                 userRepositories.UserRepository
	guidanceTemplateRepo     guidanceTemplateRepository.GuidanceTemplateRepository
	producer                 kafka_client.Producer
}

func NewIncidentService(incidentRepo repo.IncidentRepository, incidentGuidanceRepo repo.IncidentGuidanceRepository, userRepo userRepositories.UserRepository, guidanceTemplateRepo guidanceTemplateRepository.GuidanceTemplateRepository, incidentGuidanceStepRepo repo.IncidentGuidanceStepRepository, producer kafka_client.Producer) *Service {
	return &Service{incidentRepo: incidentRepo, incidentGuidanceRepo: incidentGuidanceRepo, userRepo: userRepo, guidanceTemplateRepo: guidanceTemplateRepo, incidentGuidanceStepRepo: incidentGuidanceStepRepo, producer: producer}
}

func (s *Service) CreateIncident(ctx context.Context, createIncidentDto *dto.CreateIncidentDto) (*models.Incident, error) {
	incident := &models.Incident{
		Name:        createIncidentDto.Name,
		Description: createIncidentDto.Description,
		Status:      "new",
		Severity:    createIncidentDto.Severity,
		Location:    createIncidentDto.Location,
		Alarm:       nil,
	}
	alarmID, err := uuid.Parse(createIncidentDto.AlarmId)

	if err != nil {
		return nil, errors.NewBadRequestError("Invalid asset ID format")
	}
	incident.AlarmID = alarmID

	createdIncident, err := s.incidentRepo.CreateIncident(ctx, incident)
	if err != nil {
		return nil, errors.NewDatabaseError("create incident", err)
	}

	// Validate guidance template ID
	guidanceTemplateID, err := uuid.Parse(createIncidentDto.GuidanceTemplateID)
	if err != nil {
		return nil, errors.NewBadRequestError("Invalid guidance template ID format")
	}
	guidanceTemplate, err := s.guidanceTemplateRepo.GetGuidanceTemplateByID(ctx, guidanceTemplateID.String())
	if err != nil {
		return nil, errors.NewNotFoundError("guidance template not found")
	}

	incidentGuidance := &models.IncidentGuidance{
		IncidentID: &incident.ID,
		Incident:   incident,
	}
	incidentGuidance.GuidanceTemplateID = &guidanceTemplate.ID
	incidentGuidance.GuidanceTemplate = guidanceTemplate
	// Validate guidance assignee
	assigneeId, err := uuid.Parse(createIncidentDto.Assignee)
	if err != nil {
		return nil, errors.NewBadRequestError("Invalid assignee ID format")
	}
	assigneeInfo, err := s.userRepo.GetUserByID(ctx, assigneeId.String())
	if err != nil {
		return nil, errors.NewNotFoundError("assignee not found")
	}
	incidentGuidance.AssigneeID = &assigneeInfo.ID
	incidentGuidance.Assignee = assigneeInfo

	createdIncidentGuidance, err := s.incidentGuidanceRepo.CreateIncidentGuidance(ctx, incidentGuidance)
	if err != nil {
		return nil, errors.NewDatabaseError("assign guidance", err)
	}
	steps := []models.IncidentGuidanceStep{}
	for _, step := range guidanceTemplate.GuidanceSteps {
		steps = append(steps, models.IncidentGuidanceStep{
			IncidentGuidanceID: createdIncidentGuidance.ID,
			StepNumber:         int64(step.StepNumber),
			Title:              step.Title,
			Description:        step.Description,
			IsCompleted:        false,
		})
	}

	_, _ = s.incidentGuidanceStepRepo.CreateIncidentGuidanceSteps(ctx, steps)
	// Send Kafka message
	// producerMessage := kafka.Message{
	// 	Key:   []byte(incident.ID.String()),
	// 	Value: []byte("Incident guidance assigned"),
	// }
	// if err := s.producer.WriteMessages(ctx, producerMessage); err != nil {
	// 	return nil, errors.NewAppError(errors.ErrorTypeInternal, "Failed to send Kafka message", err)
	// }

	return createdIncident, nil
}

func (s *Service) GetIncidents(ctx context.Context, page int, limit int) (*types.PaginateResponse[models.Incident], error) {
	incidents, err := s.incidentRepo.GetIncidents(ctx, page, limit)

	if err != nil {
		return nil, errors.NewDatabaseError("get incidents", err)
	}
	total, err := s.incidentRepo.GetIncidentsCount(ctx)
	totalPages := int(total) / limit
	if total%int64(limit) != 0 {
		totalPages++
	}

	if err != nil {
		return nil, errors.NewDatabaseError("get incidents count", err)
	}
	paginateResponse := types.PaginateResponse[models.Incident]{
		Pagination: types.Pagination{
			TotalPages: int(totalPages),
			Page:       page,
			Limit:      limit,
		},
		Data: incidents,
	}
	return &paginateResponse, nil
}

func (s *Service) GetIncidentByID(ctx context.Context, id string) (*models.Incident, error) {
	incident, err := s.incidentRepo.GetIncidentByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("get incident")
	}

	return incident, nil
}

func (s *Service) AssignGuidance(ctx context.Context, incidentID string, assignGuidanceDto *dto.AssignGuidance) (*models.IncidentGuidance, error) {
	incident, err := s.incidentRepo.GetIncidentByID(ctx, incidentID)
	if err != nil {
		return nil, errors.NewNotFoundError("incident not found")
	}
	if incident == nil {
		return nil, errors.NewNotFoundError("incident not found")
	}
	incidentGuidance := &models.IncidentGuidance{
		IncidentID: &incident.ID,
		Incident:   incident,
	}
	// Validate guidance template ID
	guidanceTemplateID, err := uuid.Parse(assignGuidanceDto.GuidanceTemplateID)
	if err != nil {
		return nil, errors.NewBadRequestError("Invalid guidance template ID format")
	}
	guidanceTemplate, err := s.guidanceTemplateRepo.GetGuidanceTemplateByID(ctx, guidanceTemplateID.String())
	if err != nil {
		return nil, errors.NewNotFoundError("guidance template not found")
	}
	incidentGuidance.GuidanceTemplateID = &guidanceTemplate.ID
	incidentGuidance.GuidanceTemplate = guidanceTemplate
	// Validate guidance assignee
	assigneeId, err := uuid.Parse(assignGuidanceDto.Assignee)
	if err != nil {
		return nil, errors.NewBadRequestError("Invalid assignee ID format")
	}
	assigneeInfo, err := s.userRepo.GetUserByID(ctx, assigneeId.String())
	if err != nil {
		return nil, errors.NewNotFoundError("assignee not found")
	}
	incidentGuidance.AssigneeID = &assigneeInfo.ID
	incidentGuidance.Assignee = assigneeInfo

	createdIncidentGuidance, err := s.incidentGuidanceRepo.CreateIncidentGuidance(ctx, incidentGuidance)
	if err != nil {
		return nil, errors.NewDatabaseError("assign guidance", err)
	}
	steps := []models.IncidentGuidanceStep{}
	for _, step := range guidanceTemplate.GuidanceSteps {
		steps = append(steps, models.IncidentGuidanceStep{
			IncidentGuidanceID: createdIncidentGuidance.ID,
			StepNumber:         int64(step.StepNumber),
			Title:              step.Title,
			Description:        step.Description,
			IsCompleted:        false,
		})
	}

	// _, _ = s.incidentGuidanceStepRepo.CreateIncidentGuidanceSteps(ctx, steps)
	// Send Kafka message
	producerMessage := kafka.Message{
		Key:   []byte(incident.ID.String()),
		Value: []byte("Incident guidance assigned"),
	}
	if err := s.producer.WriteMessages(ctx, producerMessage); err != nil {
		return nil, errors.NewAppError(errors.ErrorTypeInternal, "Failed to send Kafka message", err)
	}

	return createdIncidentGuidance, nil
}
func (s *Service) GetIncidentGuidance(ctx context.Context, incidentID string) (*models.IncidentGuidance, error) {
	incidentGuidance, err := s.incidentGuidanceRepo.GetIncidentGuidanceByIncidentID(ctx, incidentID)
	if err != nil {
		return nil, errors.NewDatabaseError("get incident guidance", err)
	}
	return incidentGuidance, nil
}

func (s *Service) UpdateIncident(ctx context.Context, id string, updateIncidentDto *dto.UpdateIncidentDto) (*models.Incident, error) {
	incident, err := s.incidentRepo.GetIncidentByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("incident not found")
	}
	incident.Status = updateIncidentDto.Status
	updatedIncident, err := s.incidentRepo.UpdateIncident(ctx, id, incident)
	if err != nil {
		return nil, errors.NewDatabaseError("update incident", err)
	}
	return updatedIncident, nil
}
