package repositories

import (
	"context"
	"fmt"
	"scs-operator/internal/models"

	"gorm.io/gorm"
)

type PremiseUsersRepository struct {
	db *gorm.DB
}

func NewPremiseUsersRepository(db *gorm.DB) *PremiseUsersRepository {
	return &PremiseUsersRepository{db: db}
}

func (r *PremiseUsersRepository) CreatePremiseUsers(ctx context.Context, premiseUsers []models.UserPremise) error {
	if err := r.db.WithContext(ctx).Create(premiseUsers).Error; err != nil {
		return fmt.Errorf("failed to create user premise: %w", err)
	}
	return nil
}
func (r *PremiseUsersRepository) RemovePremiseUsersByUserIds(ctx context.Context, userIds []string) error {
	if err := r.db.WithContext(ctx).Delete(&models.UserPremise{}, "user_id IN ?", userIds).Error; err != nil {
		return fmt.Errorf("failed to remove user premise: %w", err)
	}
	return nil
}
