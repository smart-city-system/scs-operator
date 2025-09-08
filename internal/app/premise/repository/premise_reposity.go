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
func (r *PremiseRepository) GetAvailableGuards(ctx context.Context, premiseID string) ([]models.User, error) {
	var guards []models.User
	if err := r.db.WithContext(ctx).Joins("JOIN guard_premises ON users.id = guard_premises.guard_id").
		Where("guard_premises.premise_id = ?", premiseID).
		Find(&guards).Error; err != nil {
		return nil, fmt.Errorf("failed to get guards: %w", err)
	}
	return guards, nil
}
