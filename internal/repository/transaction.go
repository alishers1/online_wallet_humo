package repository

import (
	"errors"
	"online_wallet_humo/pkg/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) CreateTransaction(t *models.Transaction) (uint, error) {
	if err := r.db.Create(&t).Error; err != nil {
		return 0, err
	}

	return t.ID, nil
}

func (r *TransactionRepository) TransferMoney(senderID, recipientID, serviceID uint, amount float64) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var sender, recipient models.User
	if err := tx.First(&sender, senderID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.First(&recipient, recipientID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if sender.Balance < amount {
		tx.Rollback()
		return errors.New("Sender balance has not enough money")
	}

	sender.Balance -= amount
	if err := tx.Save(&sender).Error; err != nil {
		tx.Rollback()
		return err
	}

	recipient.Balance += amount
	if err := tx.Save(&recipient).Error; err != nil {
		tx.Rollback()
		return err
	}

	transaction := models.Transaction{
		SenderID:   sender.ID,
		ReceiverID: recipient.ID,
		ServiceID:  serviceID,
		Amount:     amount,
	}
	_, err := r.CreateTransaction(&transaction)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *TransactionRepository) TransferFromCardToWallet(senderID, recipientID, cardID, serviceID uint, amount float64) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var sender, recipient models.User

	if err := tx.First(&sender, senderID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.First(&recipient, recipientID).Error; err != nil {
		tx.Rollback()
		return err
	}

	var senderCard models.Card
	if err := r.db.Where("user_id = ? AND id = ?", senderID, cardID).First(&senderCard).Error; err != nil {
		return err
	}

	if senderCard.Balance < amount {
		tx.Rollback()
		return errors.New("Sender card balance has not enough money")
	}

	senderCard.Balance -= amount
	if err := r.db.Save(&senderCard).Error; err != nil {
		tx.Rollback()
		return err
	}

	recipient.Balance += amount
	if err := tx.Save(&recipient).Error; err != nil {
		tx.Rollback()
		return err
	}

	transaction := models.Transaction{
		SenderID:   sender.ID,
		ReceiverID: recipient.ID,
		ServiceID: serviceID,
		Amount:     amount,
	}

	_, err := r.CreateTransaction(&transaction)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *TransactionRepository) GetUserTransactions(userID, offset, pageSize uint) ([]models.Transaction, error) {
	var transactions []models.Transaction

	if err := r.db.Where("sender_id = ? OR receiver_id = ?", userID, userID).
		Order("created_at DESC").
		Offset(int(offset)).
		Limit(int(pageSize)).
		Find(&transactions).Error; err != nil {
		return nil, err
		}

	return transactions, nil
}

func (r *TransactionRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
