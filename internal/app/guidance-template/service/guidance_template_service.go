package services

import (
	"context"
	guidanceStepRepositories "smart-city/internal/app/guidance-step/repository"
	"smart-city/internal/app/guidance-template/dto"
	guidanceTemplateRepositories "smart-city/internal/app/guidance-template/repository"
	"smart-city/internal/models"
	"smart-city/pkg/errors"
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
		Category:    createGuidanceTemplateDto.Category,
	}

	createdGuidanceTemplate, err := s.guidanceTemplateRepo.CreateGuidanceTemplate(ctx, guidanceTemplate)
	if err != nil {
		return nil, errors.NewDatabaseError("create guidanceTemplate", err)
	}
	steps := []models.GuidanceStep{}
	for _, step := range createGuidanceTemplateDto.Steps {
		steps = append(steps, models.GuidanceStep{
			StepNumber:         step.StepNumber,
			Instruction:        step.Instruction,
			GuidanceTemplateID: createdGuidanceTemplate.ID,
		})
	}
	if len(steps) > 0 {
		_, err = s.guidanceStepRepo.CreateGuidanceStepsByGuidanceTemplateID(ctx, steps)
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
