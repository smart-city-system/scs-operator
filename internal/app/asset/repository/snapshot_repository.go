package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

	"gorm.io/gorm"
)

type SnapshotRepository struct {
	db *gorm.DB
}

func NewSnapshotRepository(db *gorm.DB) *SnapshotRepository {
	return &SnapshotRepository{db: db}
}

func (r *SnapshotRepository) CreateSnapshot(ctx context.Context, snapshot *models.Snapshot) (*models.Snapshot, error) {
	if err := r.db.WithContext(ctx).Create(snapshot).Error; err != nil {
		return nil, fmt.Errorf("failed to create snapshot: %w", err)
	}
	return snapshot, nil
}

func (r *SnapshotRepository) GetSnapshotsByCamera(ctx context.Context, cameraID string) ([]models.Snapshot, error) {
	var snapshots []models.Snapshot
	if err := r.db.WithContext(ctx).Where("camera_id = ?", cameraID).Find(&snapshots).Error; err != nil {
		return nil, fmt.Errorf("failed to get snapshots: %w", err)
	}
	return snapshots, nil
}

func (r *SnapshotRepository) GetSnapshotByID(ctx context.Context, id string) (*models.Snapshot, error) {
	var snapshot models.Snapshot
	if err := r.db.WithContext(ctx).First(&snapshot, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get snapshot: %w", err)
	}
	return &snapshot, nil
}

func (r *SnapshotRepository) DeleteSnapshot(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Delete(&models.Snapshot{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete snapshot: %w", err)
	}
	return nil
}
