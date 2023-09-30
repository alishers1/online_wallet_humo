package service

import "online_wallet_humo/internal/repository"

type StatisticsService struct {
	staRepo *repository.StatisticsRepository
}

func NewStatisticsService(staRepo *repository.StatisticsRepository) *StatisticsService {
	return &StatisticsService{staRepo}
}

func (s *StatisticsService) GetUsersCount(startDate, endDate string) (int64, error) {
	return s.staRepo.GetUsersCount(startDate, endDate)
}

func (s *StatisticsService) GetCardsCount(startDate, endDate string) (int64, error) {
	return s.staRepo.GetCardsCount(startDate, endDate)
}

func (s *StatisticsService) GetTransactionsCount(startDate, endDate string) (int64, error) {
	return s.staRepo.GetTransactionsCount(startDate, endDate)
}
