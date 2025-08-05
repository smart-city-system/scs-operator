package services

import (
	"context"
	"smart-city/internal/app/user/dto"
	repositories "smart-city/internal/app/user/repository"
	"smart-city/internal/models"
	"smart-city/pkg/errors"
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
		return nil, errors.NewInternalError("Failed to hash password", err)
	}

	user := &models.User{
		Name:     createUserDto.Name,
		Email:    createUserDto.Email,
		Password: hashedPassword,
		Role:     createUserDto.Role,
	}

	createdUser, err := s.userRepo.CreateUser(ctx, user)
	if err != nil {
		// Check if it's a duplicate email error
		if isDuplicateEmailError(err) {
			return nil, errors.NewConflictError("User with this email already exists")
		}
		return nil, errors.NewDatabaseError("create user", err)
	}

	return createdUser, nil
}

func (s *Service) GetUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, errors.NewDatabaseError("get users", err)
	}
	return users, nil
}

// isDuplicateEmailError checks if the error is due to duplicate email constraint
func isDuplicateEmailError(err error) bool {
	errStr := err.Error()
	return contains(errStr, "duplicate key value violates unique constraint") &&
		contains(errStr, "email")
}

// contains checks if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			(len(s) > len(substr) &&
				(s[:len(substr)] == substr ||
					s[len(s)-len(substr):] == substr ||
					containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
