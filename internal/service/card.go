package service

import (
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
)

type CardService struct {
	CardRepo *repository.CardRepository
}

func NewCardService(cardRepo *repository.CardRepository) *CardService {
	return &CardService{cardRepo}
}

func (s *CardService) AddCardToUser(card *models.Card) error {
	if err := s.CardRepo.CreateCard(card); err != nil {
		return err
	}
	return nil
}

func (s *CardService) GetUserCards(userID uint) ([]models.Card, error) {
	cards, err := s.CardRepo.GetCardsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return cards, nil
}


