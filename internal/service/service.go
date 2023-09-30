package service

import (
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
)

type ServiceService struct {
	serviceRepo *repository.ServiceRepository
}

func NewServiceService(sr *repository.ServiceRepository) *ServiceService {
	return &ServiceService{sr}
}

func (service *ServiceService) CreateService(s *models.Service) (uint, error) {
	return service.serviceRepo.CreateService(s)
}

func (service *ServiceService) GetServiceByID(id uint) (*models.Service, error) {
	return service.serviceRepo.GetServiceByID(id)
}

func (service *ServiceService) UpdateServiceByID(s *models.Service) error {
	return service.serviceRepo.UpdateServiceByID(s)
}

func (service *ServiceService) DeleteServiceByID(s *models.Service) error {
	return service.serviceRepo.DeleteServiceByID(s)
}

func (service *ServiceService) GetUserByID(id uint) (*models.User, error) {
	return service.serviceRepo.GetUserByID(id)
}