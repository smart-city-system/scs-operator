package services

import (
	"context"
	guidanceStepRepositories "scs-operator/internal/app/guidance-step/repository"
	"scs-operator/internal/app/guidance-template/dto"
	guidanceTemplateRepositories "scs-operator/internal/app/guidance-template/repository"
	"scs-operator/internal/models"
	"scs-operator/pkg/errors"

	"github.com/google/uuid"
)

type Service struct {
	guidanceTemplateRepo guidanceTemplateRepositories.GuidanceTemplateRepository
	guidanceStepRepo     guidanceStepRepositories.GuidanceStepRepository
}

func NewGuidanceTemplateService(guidanceTemplateRepo guidanceTemplateRepositories.GuidanceTemplateRepository, guidanceStepRepo guidanceStepRepositories.GuidanceStepRepository) *Service {
	return &Service{guidanceTemplateRepo: guidanceTemplateRepo, guidanceStepRepo: guidanceStepRepo}
}

func (s *Service) CreateGuidanceTemplate(ctx context.Context, createGuidanceTemplateDto *dto.CreateGuidanceTemplateDto) (*models.GuidanceTemplate, error) {
	guidanceTemplate := &models.GuidanceTemplate{
		Name:        createGuidanceTemplateDto.Name,
		Description: createGuidanceTemplateDto.Description,
	}

	createdGuidanceTemplate, err := s.guidanceTemplateRepo.CreateGuidanceTemplate(ctx, guidanceTemplate)
	if err != nil {
		return nil, errors.NewDatabaseError("create guidanceTemplate", err)
	}
	steps := []models.GuidanceStep{}
	for _, step := range createGuidanceTemplateDto.Steps {
		steps = append(steps, models.GuidanceStep{
			StepNumber:         step.StepNumber,
			Title:              step.Title,
			Description:        step.Description,
			GuidanceTemplateID: createdGuidanceTemplate.ID,
		})
	}
	if len(steps) > 0 {
		_, err = s.guidanceStepRepo.CreateGuidanceSteps(ctx, steps)
		if err != nil {
			return nil, errors.NewDatabaseError("create guidanceSteps", err)
		}
	}

	return createdGuidanceTemplate, nil
}

func (s *Service) GetGuidanceTemplates(ctx context.Context) ([]models.GuidanceTemplate, error) {
	guidanceTemplates, err := s.guidanceTemplateRepo.GetGuidanceTemplates(ctx)
	if err != nil {
		return nil, errors.NewDatabaseError("get guidanceTemplates", err)
	}
	return guidanceTemplates, nil
}

func (s *Service) GetGuidanceTemplateByID(ctx context.Context, id string) (*models.GuidanceTemplate, error) {
	guidanceTemplate, err := s.guidanceTemplateRepo.GetGuidanceTemplateByID(ctx, id)
	if err != nil {
		return nil, errors.NewDatabaseError("get guidanceTemplate", err)
	}
	return guidanceTemplate, nil
}

func (s *Service) UpdateGuidanceTemplate(ctx context.Context, id string, updateGuidanceTemplateDto *dto.UpdateGuidanceTemplateDto) (*models.GuidanceTemplate, error) {
	guidanceTemplate, err := s.guidanceTemplateRepo.GetGuidanceTemplateByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("guidance template not found")
	}
	guidanceTemplate.Name = updateGuidanceTemplateDto.Name
	guidanceTemplate.Description = updateGuidanceTemplateDto.Description
	updatedGuidanceTemplate, err := s.guidanceTemplateRepo.UpdateGuidanceTemplate(ctx, id, guidanceTemplate)
	if err != nil {
		return nil, errors.NewDatabaseError("update guidance template", err)
	}
	if updateGuidanceTemplateDto.AddSteps != nil {
		steps := []models.GuidanceStep{}
		for _, step := range updateGuidanceTemplateDto.AddSteps {
			steps = append(steps, models.GuidanceStep{
				StepNumber:         step.StepNumber,
				Title:              step.Title,
				Description:        step.Description,
				GuidanceTemplateID: updatedGuidanceTemplate.ID,
			})
		}
		if len(steps) > 0 {
			_, err = s.guidanceStepRepo.CreateGuidanceSteps(ctx, steps)
			if err != nil {
				return nil, errors.NewDatabaseError("create guidanceSteps", err)
			}
		}
		if updateGuidanceTemplateDto.UpdateSteps != nil {
			for _, step := range updateGuidanceTemplateDto.UpdateSteps {
				if step.ID == nil {
					continue
				}
				id, err := uuid.Parse(*step.ID)
				if err != nil {
					return nil, errors.NewBadRequestError("Invalid step ID format")
				}
				err = s.guidanceStepRepo.UpdateGuidanceStep(ctx, id.String(), &models.GuidanceStep{
					StepNumber:  step.StepNumber,
					Title:       step.Title,
					Description: step.Description,
				})
				if err != nil {
					return nil, errors.NewDatabaseError("update guidanceStep", err)
				}
			}
			if updateGuidanceTemplateDto.RemoveSteps != nil {
				removeStepIds := []string{}
				for _, stepID := range updateGuidanceTemplateDto.RemoveSteps {
					id, err := uuid.Parse(stepID)
					if err != nil {
						return nil, errors.NewBadRequestError("Invalid step ID format")
					}
					removeStepIds = append(removeStepIds, id.String())
				}
				if len(removeStepIds) > 0 {
					err = s.guidanceStepRepo.DeleteGuidanceSteps(ctx, removeStepIds)
					if err != nil {
						return nil, errors.NewDatabaseError("delete guidanceSteps", err)
					}
				}
			}
		}
	}
	return updatedGuidanceTemplate, nil
}
