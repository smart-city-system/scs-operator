package repositories

import (
	"context"
	"fmt"
	"scs-operator/internal/models"

	"gorm.io/gorm"
)

type PremiseRepository struct {
	db *gorm.DB
}

func NewPremiseRepository(db *gorm.DB) *PremiseRepository {
	return &PremiseRepository{db: db}
}

func (r *PremiseRepository) CreatePremise(ctx context.Context, Premise *models.Premise) (*models.Premise, error) {
	if err := r.db.WithContext(ctx).Create(Premise).Error; err != nil {
		return nil, fmt.Errorf("failed to create Premise: %w", err)
	}
	return Premise, nil
}

func (r *PremiseRepository) GetPremises(ctx context.Context, page int, limit int) ([]models.Premise, error) {
	var Premises []models.Premise
	if err := r.db.WithContext(ctx).Limit(limit).Offset((page - 1) * limit).Find(&Premises).Error; err != nil {
		return nil, fmt.Errorf("failed to get Premises: %w", err)
	}
	return Premises, nil
}

func (r *PremiseRepository) GetPremisesCount(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.Premise{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to get Premises count: %w", err)
	}
	return count, nil
}

func (r *PremiseRepository) GetPremiseByID(ctx context.Context, id string) (*models.Premise, error) {
	var Premise models.Premise

	if err := r.db.WithContext(ctx).First(&Premise, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get Premise: %w", err)
	}

	return &Premise, nil
}
func (r *PremiseRepository) GetAvailableUsers(ctx context.Context, premiseID string) ([]models.User, error) {
	var users []models.User
	if err := r.db.WithContext(ctx).Joins("JOIN user_premises ON users.id = user_premises.user_id").
		Where("user_premises.premise_id = ?", premiseID).
		Find(&users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return users, nil
}

func (r *PremiseRepository) UpdatePremise(ctx context.Context, id string, premise *models.Premise) (*models.Premise, error) {
	result := r.db.WithContext(ctx).Model(&models.Premise{}).Where("id = ?", id).Updates(premise)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to update premise: %w", result.Error)
	}
	return premise, nil
}
