package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID        uint
	User          User
	TransactionID uint `json:"transaction_id" gorm:"column:transaction_id"`
	Transaction   Transaction `json:"-"`
}
