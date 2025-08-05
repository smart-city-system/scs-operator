package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

	"gorm.io/gorm"
)

type GuidanceStepRepository struct {
	db *gorm.DB
}

func NewGuidanceStepRepository(db *gorm.DB) *GuidanceStepRepository {
	return &GuidanceStepRepository{db: db}
}

func (r *GuidanceStepRepository) CreateGuidanceStep(ctx context.Context, GuidanceStep *models.GuidanceStep) (*models.GuidanceStep, error) {
	if err := r.db.WithContext(ctx).Create(GuidanceStep).Error; err != nil {
		return nil, fmt.Errorf("failed to create GuidanceStep: %w", err)
	}
	return GuidanceStep, nil
}
func (r *GuidanceStepRepository) GetGuidanceSteps(ctx context.Context) ([]models.GuidanceStep, error) {
	var GuidanceSteps []models.GuidanceStep
	if err := r.db.WithContext(ctx).Find(&GuidanceSteps).Error; err != nil {
		return nil, fmt.Errorf("failed to get GuidanceSteps: %w", err)
	}
	return GuidanceSteps, nil
}

func (r *GuidanceStepRepository) GetGuidanceStepByID(ctx context.Context, id string) (*models.GuidanceStep, error) {
	var GuidanceStep models.GuidanceStep
	if err := r.db.WithContext(ctx).First(&GuidanceStep, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get GuidanceStep: %w", err)
	}
	return &GuidanceStep, nil
}

func (r *GuidanceStepRepository) CreateGuidanceStepsByGuidanceTemplateID(ctx context.Context, steps []models.GuidanceStep) ([]models.GuidanceStep, error) {
	if err := r.db.WithContext(ctx).Create(&steps).Error; err != nil {
		return nil, fmt.Errorf("failed to create GuidanceSteps: %w", err)
	}
	return steps, nil
}
