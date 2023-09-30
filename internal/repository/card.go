package repository

import (
	"online_wallet_humo/pkg/models"
	"gorm.io/gorm"
)

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db}
}

func (r *CardRepository) CreateCard(card *models.Card) error {
	if err := r.db.Create(&card).Error; err != nil {
		return err
	}

	return nil
}

func (r *CardRepository) GetCardsByUserID(userID uint) ([]models.Card, error) {
	var cards []models.Card

	if err := r.db.Where("user_id = ?", userID).Find(&cards).Error; err != nil {
		return nil, err
	}

	return cards, nil
}

func (r *CardRepository) GetCardByID(cardID uint) (*models.Card, error) {
	var card models.Card

	if err := r.db.Where("id = ?", cardID).Find(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

