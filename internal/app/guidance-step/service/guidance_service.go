package services

import (
	"context"
	"scs-operator/internal/app/guidance-step/dto"
	repositories "scs-operator/internal/app/guidance-step/repository"
	"scs-operator/internal/models"
	"scs-operator/pkg/errors"

	"github.com/google/uuid"
)

type Service struct {
	guidanceStepRepo repositories.GuidanceStepRepository
}

func NewGuidanceStepService(guidanceStepRepo repositories.GuidanceStepRepository) *Service {
	return &Service{guidanceStepRepo: guidanceStepRepo}
}

func (s *Service) CreateGuidanceStep(ctx context.Context, createGuidanceStepDto *dto.CreateGuidanceStepDto) (*models.GuidanceStep, error) {

	guidanceStep := &models.GuidanceStep{
		StepNumber:  createGuidanceStepDto.StepNumber,
		Title:       createGuidanceStepDto.Title,
		Description: createGuidanceStepDto.Description,
	}
	guidanceTemplateID, err := uuid.Parse(createGuidanceStepDto.GuidanceTemplateID)

	if err != nil {
		return nil, errors.NewBadRequestError("Invalid guidanceTemplate ID format")
	}
	guidanceStep.GuidanceTemplateID = guidanceTemplateID

	createdGuidanceStep, err := s.guidanceStepRepo.CreateGuidanceStep(ctx, guidanceStep)
	if err != nil {
		return nil, errors.NewDatabaseError("create guidanceStep", err)
	}

	return createdGuidanceStep, nil
}

func (s *Service) GetGuidanceSteps(ctx context.Context) ([]models.GuidanceStep, error) {
	guidanceSteps, err := s.guidanceStepRepo.GetGuidanceSteps(ctx)
	if err != nil {
		return nil, errors.NewDatabaseError("get guidanceSteps", err)
	}
	return guidanceSteps, nil
}

func (s *Service) GetGuidanceStepByID(ctx context.Context, id string) (*models.GuidanceStep, error) {
	guidanceStep, err := s.guidanceStepRepo.GetGuidanceStepByID(ctx, id)
	if err != nil {
		return nil, errors.NewDatabaseError("get guidanceStep", err)
	}
	return guidanceStep, nil
}
