package services

import (
	"context"
	"scs-operator/internal/models"
	"scs-operator/pkg/utils"
)

type guardRepository interface {
	Create(ctx context.Context, guard *models.User) (*models.User, error)
	GetGuards(ctx context.Context) ([]models.User, error)
	Update(ctx context.Context, guard *models.User) error
	Delete(ctx context.Context, guard *models.User) error
}
type Service struct {
	guardRepo guardRepository
}

func NewGuardService(guardRepo guardRepository) *Service {
	return &Service{guardRepo: guardRepo}
}

func (s *Service) Create(ctx context.Context, guard *models.User) (*models.User, error) {
	// Hash the password before saving
	if guard.Password != "" {
		hashedPassword, err := utils.HashPassword(guard.Password)
		if err != nil {
			return nil, err
		}
		guard.Password = hashedPassword
	}

	return s.guardRepo.Create(ctx, guard)
}

func (s *Service) GetGuards(ctx context.Context) ([]models.User, error) {
	return s.guardRepo.GetGuards(ctx)
}
