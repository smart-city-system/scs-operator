package repositories

import (
	"context"
	"fmt"
	"scs-operator/internal/models"

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
func (r *UserRepository) GetUsers(ctx context.Context, page int, limit int) ([]models.User, error) {
	var Users []models.User
	if err := r.db.WithContext(ctx).Limit(limit).Offset((page - 1) * limit).Find(&Users).Error; err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	return Users, nil
}

func (r *UserRepository) GetUsersCount(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to get users count: %w", err)
	}
	return count, nil
}
func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var User models.User
	if err := r.db.WithContext(ctx).First(&User, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &User, nil
}
