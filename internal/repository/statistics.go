package repository

import (
	"online_wallet_humo/pkg/models"

	"gorm.io/gorm"
)

type StatisticsRepository struct {
	db *gorm.DB
}

func NewStatisticsRepository(db *gorm.DB) *StatisticsRepository {
	return &StatisticsRepository{db}
}

func (r *StatisticsRepository) GetUsersCount(startDate, endDate string) (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Count(&count).Error 
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *StatisticsRepository) GetCardsCount(startDate, endDate string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Card{}).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *StatisticsRepository) GetTransactionsCount(startDate, endDate string) (int64, error) {
	var count int64
	err := r.db.Model(&models.Transaction{}).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}



