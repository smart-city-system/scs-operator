package repositories

import (
	"context"
	"fmt"
	"scs-operator/internal/models"

	"gorm.io/gorm"
)

type GuardRepository struct {
	db *gorm.DB
}

func NewGuardRepository(db *gorm.DB) *GuardRepository {
	return &GuardRepository{db: db}
}

func (r *GuardRepository) Create(ctx context.Context, guard *models.User) (*models.User, error) {
	if err := r.db.WithContext(ctx).Create(guard).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return guard, nil
}
func (r *GuardRepository) GetGuards(ctx context.Context, page int, limit int) ([]models.User, error) {
	var guards []models.User
	if err := r.db.WithContext(ctx).Limit(limit).Offset((page - 1) * limit).Where("role = 'guard'").Find(&guards).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return guards, nil
}

func (r *GuardRepository) GetGuardsCount(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to get guards count: %w", err)
	}
	return count, nil
}
