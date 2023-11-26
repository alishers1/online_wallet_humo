package service

import (
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
)

type FavoriteService struct {
	favRepo *repository.FavoriteRepository
}

func NewFavoriteService(favRepo *repository.FavoriteRepository) *FavoriteService {
	return &FavoriteService{favRepo}
}

func (s *FavoriteService) CreateFavorite(favorite models.Favorite) error {
	return s.favRepo.CreateFavorite(favorite)
}

func (s *FavoriteService) GetUserFavoriteTransactions(userID, page, pageSize uint) ([]models.Favorite, error) {
	offset := (page - 1) * pageSize

	return s.favRepo.GetUserFavoriteTransactions(userID, offset, pageSize)
}