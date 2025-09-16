package services

import (
	"context"
	"scs-operator/internal/app/premise/dto"
	repositories "scs-operator/internal/app/premise/repository"
	"scs-operator/internal/models"
	"scs-operator/internal/types"
	"scs-operator/pkg/errors"

	"github.com/google/uuid"
)

type Service struct {
	premiseRepo      repositories.PremiseRepository
	premiseUsersRepo repositories.PremiseUsersRepository
}

func NewPremiseService(premiseRepo repositories.PremiseRepository, premiseUsersRepo repositories.PremiseUsersRepository) *Service {
	return &Service{premiseRepo: premiseRepo, premiseUsersRepo: premiseUsersRepo}
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

func (s *Service) GetPremises(ctx context.Context, page int, limit int) (*types.PaginateResponse[models.Premise], error) {
	premises, err := s.premiseRepo.GetPremises(ctx, page, limit)
	if err != nil {
		return nil, errors.NewDatabaseError("get premises", err)
	}
	total, err := s.premiseRepo.GetPremisesCount(ctx)
	totalPages := int(total) / limit
	if total%int64(limit) != 0 {
		totalPages++
	}

	if err != nil {
		return nil, errors.NewDatabaseError("get premises count", err)
	}
	paginateResponse := &types.PaginateResponse[models.Premise]{
		Pagination: types.Pagination{
			TotalPages: int(totalPages),
			Page:       page,
			Limit:      limit,
		},
		Data: premises,
	}
	return paginateResponse, nil
}
func (s *Service) GetPremiseByID(ctx context.Context, id string) (*models.Premise, error) {
	premise, err := s.premiseRepo.GetPremiseByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("premise")
	}
	return premise, nil
}

func (s *Service) GetAvailableUsers(ctx context.Context, premiseID string) ([]models.User, error) {
	guards, err := s.premiseRepo.GetAvailableUsers(ctx, premiseID)
	if err != nil {
		return nil, errors.NewDatabaseError("get available guards", err)
	}
	return guards, nil
}
func (s *Service) UpdatePremise(ctx context.Context, id string, updatePremiseDto *dto.UpdatePremiseDto) (*models.Premise, error) {
	premise, err := s.premiseRepo.GetPremiseByID(ctx, id)
	if err != nil {
		return nil, errors.NewNotFoundError("premise")
	}
	premise.Name = updatePremiseDto.Name
	premise.Address = updatePremiseDto.Address
	updatedPremise, err := s.premiseRepo.UpdatePremise(ctx, id, premise)
	if err != nil {
		return nil, errors.NewDatabaseError("update premise", err)
	}
	return updatedPremise, nil
}
func (s *Service) AssignUsers(ctx context.Context, premiseID string, updatePremiseUserDto *dto.UpdatePremiseUserDto) error {
	premise, err := s.premiseRepo.GetPremiseByID(ctx, premiseID)
	if err != nil {
		return errors.NewNotFoundError("premise")
	}

	// Get added users
	addedUsers := []models.UserPremise{}
	for _, userID := range updatePremiseUserDto.AddedUsers {
		userID, err := uuid.Parse(userID)
		if err != nil {
			return errors.NewBadRequestError("Invalid user ID format")
		}
		addedUser := models.UserPremise{
			UserID:    userID,
			PremiseID: premise.ID,
		}
		addedUsers = append(addedUsers, addedUser)

	}
	// Get removed users
	removedUsers := []string{}
	for _, userID := range updatePremiseUserDto.RemovedUsers {
		userID, err := uuid.Parse(userID)
		if err != nil {
			return errors.NewBadRequestError("Invalid user ID format")
		}
		removedUsers = append(removedUsers, userID.String())
	}
	if len(addedUsers) > 0 {
		err = s.premiseUsersRepo.CreatePremiseUsers(ctx, addedUsers)
		if err != nil {
			return errors.NewDatabaseError("add premise users", err)
		}
	}
	if (len(removedUsers) > 0) && (removedUsers[0] != "") {
		err = s.premiseUsersRepo.RemovePremiseUsersByUserIds(ctx, removedUsers)
		if err != nil {
			return errors.NewDatabaseError("remove premise users", err)
		}
	}
	return nil
}
