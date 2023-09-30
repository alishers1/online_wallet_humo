package repository

import (
	"online_wallet_humo/pkg/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (p *AuthPostgres) CreateUser(u models.User) (uint, error) {
	if err := p.db.Create(&u).Error; err != nil {
		return 0, err
	}

	return u.ID, nil
}

func (p *AuthPostgres) GetUser(email string) (*models.User, error) {
	var u models.User
	if err := p.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}

	return &u, nil
}