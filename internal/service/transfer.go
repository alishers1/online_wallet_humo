package service

import (
	"online_wallet_humo/internal/repository"
	"online_wallet_humo/pkg/models"
)

type TransferService struct {
	transactionRepo *repository.TransactionRepository
}

func NewTransferService(tr *repository.TransactionRepository) *TransferService {
	return &TransferService{tr}
}

func (s *TransferService) TransferFromWalletToWallet(senderID, recipientID, serviceID uint, amount float64) error {
	return s.transactionRepo.TransferMoney(senderID, recipientID, serviceID, amount)
}

func (s *TransferService) TransferFromCardToWallet(senderID, recipientID, cardID uint, amount float64) error {
	return s.transactionRepo.TransferFromCardToWallet(senderID, recipientID, cardID, amount)
}

func (s *TransferService) GetUserTransactions(userID, page, pageSize uint) ([]models.Transaction, error) {
	offset := (page - 1) * pageSize

	transactions, err := s.transactionRepo.GetUserTransactions(userID, offset, pageSize)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}