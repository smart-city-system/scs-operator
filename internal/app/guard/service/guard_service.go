package services

import (
	"context"
	"scs-operator/internal/app/guard/dto"
	repositories "scs-operator/internal/app/guard/repository"
	"scs-operator/internal/models"
	"scs-operator/pkg/utils"

	"github.com/google/uuid"
)

type Service struct {
	guardRepo        repositories.GuardRepository
	guardPremiseRepo repositories.GuardPremiseRepository
}

func NewGuardService(guardRepo repositories.GuardRepository, guardPremiseRepo repositories.GuardPremiseRepository) *Service {
	return &Service{guardRepo: guardRepo, guardPremiseRepo: guardPremiseRepo}
}

func (s *Service) Create(ctx context.Context, createGuardDto *dto.CreateGuardDto) (*models.User, error) {
	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(createGuardDto.Password)
	if err != nil {
		return nil, err
	}
	guard := &models.User{
		Name:     createGuardDto.Name,
		Email:    createGuardDto.Email,
		Password: hashedPassword,
		Role:     "guard",
	}

	createdGuard, err := s.guardRepo.Create(ctx, guard)
	if err != nil {
		return nil, err
	}
	return createdGuard, nil

}

func (s *Service) GetGuards(ctx context.Context) ([]models.User, error) {
	return s.guardRepo.GetGuards(ctx)
}

func (s *Service) AssignPremises(ctx context.Context, guardID string, premiseID string) error {
	guard, err := uuid.Parse(guardID)
	if err != nil {
		return err
	}
	premise, err := uuid.Parse(premiseID)
	if err != nil {
		return err
	}
	guardPremise := &models.GuardPremise{
		GuardID:   guard,
		PremiseID: premise,
	}
	// Check if the guard is already assigned to the premise
	exists, err := s.guardPremiseRepo.CheckExist(ctx, guardPremise)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	_, err = s.guardPremiseRepo.AssignPremises(ctx, guardPremise)
	if err != nil {
		return err
	}
	return nil
}
