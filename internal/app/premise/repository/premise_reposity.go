package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

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
func (r *PremiseRepository) GetPremises(ctx context.Context) ([]models.Premise, error) {
	var Premises []models.Premise
	if err := r.db.WithContext(ctx).Find(&Premises).Error; err != nil {
		return nil, fmt.Errorf("failed to get Premises: %w", err)
	}
	return Premises, nil
}

func (r *PremiseRepository) GetPremiseByID(ctx context.Context, id string) (*models.Premise, error) {
	var Premise models.Premise

	if err := r.db.WithContext(ctx).First(&Premise, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get Premise: %w", err)
	}

	return &Premise, nil
}
