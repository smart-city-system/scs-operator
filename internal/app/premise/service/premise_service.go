package services

import (
	"context"
	"scs-operator/internal/app/premise/dto"
	repositories "scs-operator/internal/app/premise/repository"
	"scs-operator/internal/models"
	"scs-operator/pkg/errors"

	"github.com/google/uuid"
)

type Service struct {
	premiseRepo repositories.PremiseRepository
}

func NewPremiseService(premiseRepo repositories.PremiseRepository) *Service {
	return &Service{premiseRepo: premiseRepo}
}

func (s *Service) CreatePremise(ctx context.Context, createPremiseDto *dto.CreatePremiseDto) (*models.Premise, error) {
	premise := &models.Premise{
		Name:    createPremiseDto.Name,
		Address: createPremiseDto.Address,
	}
	if createPremiseDto.ParentPremiseID != "" {
		parentID, err := uuid.Parse(createPremiseDto.ParentPremiseID)
		if err != nil {
			return nil, errors.NewBadRequestError("Invalid parent premise ID format")
		}
		// Verify that the parent premise exists
		_, err = s.premiseRepo.GetPremiseByID(ctx, parentID.String())
		if err != nil {
			return nil, errors.NewNotFoundError("parent premise")
		}
		// Set the ParentPremiseID for the foreign key relationship
		premise.ParentPremiseID = &parentID
	}

	createdPremise, err := s.premiseRepo.CreatePremise(ctx, premise)
	if err != nil {
		return nil, errors.NewDatabaseError("create premise", err)
	}
	return createdPremise, nil
}

func (s *Service) GetPremises(ctx context.Context) ([]models.Premise, error) {
	premises, err := s.premiseRepo.GetPremises(ctx)
	if err != nil {
		return nil, errors.NewDatabaseError("get premises", err)
	}
	return premises, nil
}
