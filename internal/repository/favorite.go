package repository

import (
	"online_wallet_humo/pkg/models"

	"gorm.io/gorm"
)

type FavoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db}
}

func (r *FavoriteRepository) CreateFavorite(favorite models.Favorite) error {
	if err := r.db.Create(&favorite).Error; err != nil {
		return err
	}

	return nil
}

func (r *FavoriteRepository) GetUserFavoriteTransactions(userID, offset, pageSize uint) ([]models.Favorite, error) {
	var favTransactions []models.Favorite

	if err := r.db.Where("user_id", userID).Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&favTransactions).Error; err != nil {
			return nil, err
		}

	return favTransactions, nil
}