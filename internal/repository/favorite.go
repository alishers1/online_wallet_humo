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