package services

import (
	"context"
	incidentRepositories "scs-operator/internal/app/incident/repository"
	"scs-operator/internal/app/user/dto"
	repositories "scs-operator/internal/app/user/repository"
	"scs-operator/internal/models"
	"scs-operator/internal/types"
	"scs-operator/pkg/errors"
	"scs-operator/pkg/utils"
)

type Service struct {
	userRepo                 repositories.UserRepository
	incidentGuidanceRepo     incidentRepositories.IncidentGuidanceRepository
	incidentGuidanceStepRepo incidentRepositories.IncidentGuidanceStepRepository
}

func NewUserService(userRepo repositories.UserRepository, incidentGuidanceRepo incidentRepositories.IncidentGuidanceRepository, incidentGuidanceStepRepo incidentRepositories.IncidentGuidanceStepRepository) *Service {
	return &Service{userRepo: userRepo, incidentGuidanceRepo: incidentGuidanceRepo, incidentGuidanceStepRepo: incidentGuidanceStepRepo}
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

func (s *Service) GetUsers(ctx context.Context, page int, limit int) (*types.PaginateResponse[models.User], error) {
	users, err := s.userRepo.GetUsers(ctx, page, limit)
	if err != nil {
		return nil, errors.NewDatabaseError("get users", err)
	}
	total, err := s.userRepo.GetUsersCount(ctx)
	totalPages := int(total) / limit
	if total%int64(limit) != 0 {
		totalPages++
	}

	if err != nil {
		return nil, errors.NewDatabaseError("get users count", err)
	}
	paginateResponse := &types.PaginateResponse[models.User]{
		Pagination: types.Pagination{
			TotalPages: int(totalPages),
			Page:       page,
			Limit:      limit,
		},
		Data: users,
	}
	return paginateResponse, nil
}

func (s *Service) GetAssignments(ctx context.Context, userID string) ([]models.IncidentGuidance, error) {
	assignments, err := s.incidentGuidanceRepo.GetIncidentGuidanceByAssigneeID(ctx, userID)
	if err != nil {
		return nil, errors.NewDatabaseError("get assignments", err)
	}
	return assignments, nil
}

func (s *Service) CompleteStep(ctx context.Context, assignmentID string, stepID string) error {
	// TODO: Check if the step belongs to the assignment
	stepInfo, err := s.incidentGuidanceStepRepo.GetIncidentGuidanceStepByID(ctx, stepID)
	if err != nil {
		return errors.NewDatabaseError("get step", err)
	}
	if stepInfo.IncidentGuidanceID.String() != assignmentID {
		return errors.NewBadRequestError("step does not belong to the assignment")
	}
	if stepInfo.IsCompleted {
		return errors.NewBadRequestError("step already completed")
	}

	s.incidentGuidanceStepRepo.UpdateIncidentGuidanceStep(ctx, stepID, true)
	return nil
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
