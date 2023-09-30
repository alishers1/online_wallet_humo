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