package repositories

import (
	"context"
	"fmt"
	"scs-operator/internal/models"

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

func (r *GuidanceStepRepository) CreateGuidanceSteps(ctx context.Context, steps []models.GuidanceStep) ([]models.GuidanceStep, error) {
	if err := r.db.WithContext(ctx).Create(steps).Error; err != nil {
		return nil, fmt.Errorf("failed to create GuidanceSteps: %w", err)
	}
	return steps, nil
}

func (r *GuidanceStepRepository) UpdateGuidanceStep(ctx context.Context, id string, guidanceStep *models.GuidanceStep) error {
	result := r.db.WithContext(ctx).Model(&models.GuidanceStep{}).Where("id = ?", id).Updates(guidanceStep)
	if result.Error != nil {
		return fmt.Errorf("failed to update guidance step: %w", result.Error)
	}
	return nil
}
func (r *GuidanceStepRepository) DeleteGuidanceSteps(ctx context.Context, ids []string) error {
	if err := r.db.WithContext(ctx).Delete(&models.GuidanceStep{}, "id IN ?", ids).Error; err != nil {
		return fmt.Errorf("failed to delete guidance steps: %w", err)
	}
	return nil
}
