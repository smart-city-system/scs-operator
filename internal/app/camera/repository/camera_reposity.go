package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

	"gorm.io/gorm"
)

type CameraRepository struct {
	db *gorm.DB
}

func NewCameraRepository(db *gorm.DB) *CameraRepository {
	return &CameraRepository{db: db}
}

func (r *CameraRepository) CreateCamera(ctx context.Context, Camera *models.Camera) (*models.Camera, error) {
	if err := r.db.WithContext(ctx).Create(Camera).Error; err != nil {
		return nil, fmt.Errorf("failed to create Camera: %w", err)
	}
	return Camera, nil
}
func (r *CameraRepository) GetCameras(ctx context.Context) ([]models.Camera, error) {
	var Cameras []models.Camera
	if err := r.db.WithContext(ctx).Find(&Cameras).Error; err != nil {
		return nil, fmt.Errorf("failed to get Cameras: %w", err)
	}
	return Cameras, nil
}

func (r *CameraRepository) GetCameraByID(ctx context.Context, id string) (*models.Camera, error) {
	var Camera models.Camera
	if err := r.db.WithContext(ctx).First(&Camera, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get Camera: %w", err)
	}
	return &Camera, nil
}
