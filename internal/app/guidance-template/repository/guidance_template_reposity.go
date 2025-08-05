package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

	"gorm.io/gorm"
)

type GuidanceTemplateRepository struct {
	db *gorm.DB
}

func NewGuidanceTemplateRepository(db *gorm.DB) *GuidanceTemplateRepository {
	return &GuidanceTemplateRepository{db: db}
}

func (r *GuidanceTemplateRepository) CreateGuidanceTemplate(ctx context.Context, GuidanceTemplate *models.GuidanceTemplate) (*models.GuidanceTemplate, error) {
	if err := r.db.WithContext(ctx).Create(GuidanceTemplate).Error; err != nil {
		return nil, fmt.Errorf("failed to create GuidanceTemplate: %w", err)
	}
	return GuidanceTemplate, nil
}
func (r *GuidanceTemplateRepository) GetGuidanceTemplates(ctx context.Context) ([]models.GuidanceTemplate, error) {
	var GuidanceTemplates []models.GuidanceTemplate
	if err := r.db.WithContext(ctx).Preload("GuidanceSteps").Find(&GuidanceTemplates).Error; err != nil {
		return nil, fmt.Errorf("failed to get GuidanceTemplates: %w", err)
	}
	return GuidanceTemplates, nil
}

func (r *GuidanceTemplateRepository) GetGuidanceTemplateByID(ctx context.Context, id string) (*models.GuidanceTemplate, error) {
	var GuidanceTemplate models.GuidanceTemplate
	if err := r.db.WithContext(ctx).First(&GuidanceTemplate, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get GuidanceTemplate: %w", err)
	}
	return &GuidanceTemplate, nil
}
