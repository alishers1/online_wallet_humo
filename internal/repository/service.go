package repository

import (
	"online_wallet_humo/pkg/models"

	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db}
}

func (r *ServiceRepository) CreateService(s *models.Service) (uint, error) {
	if err := r.db.Create(&s).Error; err != nil {
		return 0, err
	}

	return s.ID, nil
}

func (r *ServiceRepository) GetServiceByID(id uint) (*models.Service, error) {
	var service models.Service

	if err := r.db.Where("id = ?", id).First(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

func (r *ServiceRepository) UpdateServiceByID(s *models.Service) error {
	if err := r.db.Where("id = ?", s.ID).Save(s).Error; err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) DeleteServiceByID(serviceID uint) error {
	var service models.Service

	if err := r.db.First(&service, serviceID).Error; err != nil {
		return err
	}

	service.IsActive = false

	if err := r.db.Save(&service).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&service).Error; err != nil {
		return err
	}

	return nil
}

func (r *ServiceRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
