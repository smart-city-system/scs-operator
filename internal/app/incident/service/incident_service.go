package services

import (
	"context"
	"smart-city/internal/app/incident/dto"
	repositories "smart-city/internal/app/incident/repository"
	"smart-city/internal/models"
	"smart-city/pkg/errors"

	"github.com/google/uuid"
)

type Service struct {
	incidentRepo repositories.IncidentRepository
}

func NewIncidentService(incidentRepo repositories.IncidentRepository) *Service {
	return &Service{incidentRepo: incidentRepo}
}

func (s *Service) CreateIncident(ctx context.Context, createIncidentDto *dto.CreateIncidentDto) (*models.Incident, error) {

	incident := &models.Incident{
		Name:        createIncidentDto.Name,
		Description: createIncidentDto.Description,
		Location:    createIncidentDto.Location,
		Status:      createIncidentDto.Status,
		AlertID:     nil,
		Asset:       nil,
	}
	// alertID, err := uuid.Parse(createIncidentDto.AlertID)
	// if err != nil {
	// 	return nil, errors.NewBadRequestError("Invalid alert ID format")
	// }
	// incident.AlertID = &alertID
	assetID, err := uuid.Parse(createIncidentDto.AssetID)

	if err != nil {
		return nil, errors.NewBadRequestError("Invalid asset ID format")
	}
	incident.AssetID = assetID

	createdIncident, err := s.incidentRepo.CreateIncident(ctx, incident)
	if err != nil {
		return nil, errors.NewDatabaseError("create incident", err)
	}

	return createdIncident, nil
}

func (s *Service) GetIncidents(ctx context.Context) ([]models.Incident, error) {
	incidents, err := s.incidentRepo.GetIncidents(ctx)
	if err != nil {
		return nil, errors.NewDatabaseError("get incidents", err)
	}
	return incidents, nil
}

func (s *Service) GetIncidentByID(ctx context.Context, id string) (*models.Incident, error) {
	incident, err := s.incidentRepo.GetIncidentByID(ctx, id)
	if err != nil {
		return nil, errors.NewDatabaseError("get incident", err)
	}
	return incident, nil
}
