package repositories

import (
	"context"
	"fmt"
	"smart-city/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, User *models.User) (*models.User, error) {
	if err := r.db.WithContext(ctx).Create(User).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return User, nil
}
func (r *UserRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	var Users []models.User
	if err := r.db.WithContext(ctx).Find(&Users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return Users, nil
}
