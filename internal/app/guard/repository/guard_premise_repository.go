package repositories

import (
	"gorm.io/gorm"
)

type GuardPremiseRepository struct {
	db *gorm.DB
}

func NewGuardPremiseRepository(db *gorm.DB) *GuardPremiseRepository {
	return &GuardPremiseRepository{db: db}
}

// func (r *GuardPremiseRepository) AssignPremises(ctx context.Context, guardPremise *models.GuardPremise) (*models.GuardPremise, error) {
// 	if err := r.db.WithContext(ctx).Create(guardPremise).Error; err != nil {
// 		return nil, fmt.Errorf("failed to create user: %w", err)
// 	}
// 	return guardPremise, nil
// }

// func (r *GuardPremiseRepository) CheckExist(ctx context.Context, guardPremise *models.GuardPremise) (bool, error) {
// 	existingGuardPremise := &models.GuardPremise{}
// 	if err := r.db.WithContext(ctx).Where("guard_id = ? AND premise_id = ?", guardPremise.GuardID, guardPremise.PremiseID).First(existingGuardPremise).Error; err == nil {
// 		return true, nil
// 	}
// 	return false, nil
// }
