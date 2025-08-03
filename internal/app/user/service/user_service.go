package services

import (
	"context"
	"smart-city/internal/app/user/dto"
	repositories "smart-city/internal/app/user/repository"
	"smart-city/internal/models"
	"smart-city/pkg/utils"
)

type Service struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *Service {
	return &Service{userRepo: userRepo}
}

func (s *Service) CreateUser(ctx context.Context, createUserDto *dto.CreateUserDto) (*models.User, error) {
	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(createUserDto.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Password: hashedPassword,
		Role:     createUserDto.Role,
	}

	return s.userRepo.CreateUser(ctx, user)
}

func (s *Service) GetUsers(ctx context.Context) ([]models.User, error) {
	return s.userRepo.GetUsers(ctx)
}
