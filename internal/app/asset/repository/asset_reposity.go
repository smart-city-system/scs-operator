package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

	"gorm.io/gorm"
)

type AssetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) CreateAsset(ctx context.Context, Asset *models.Asset) (*models.Asset, error) {
	if err := r.db.WithContext(ctx).Create(Asset).Error; err != nil {
		return nil, fmt.Errorf("failed to create Asset: %w", err)
	}
	return Asset, nil
}
func (r *AssetRepository) GetAssets(ctx context.Context) ([]models.Asset, error) {
	var Assets []models.Asset
	if err := r.db.WithContext(ctx).Find(&Assets).Error; err != nil {
		return nil, fmt.Errorf("failed to get Assets: %w", err)
	}
	return Assets, nil
}

func (r *AssetRepository) GetAssetByID(ctx context.Context, id string) (*models.Asset, error) {
	var Asset models.Asset
	if err := r.db.WithContext(ctx).First(&Asset, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get Asset: %w", err)
	}
	return &Asset, nil
}
