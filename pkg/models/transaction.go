package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	SenderID   uint    `json:"sender_id" gorm:"column:sender_id"`
	Sender     User    `json:"-"`
	ReceiverID uint    `json:"receiver_id" gorm:"column:receiver_id"`
	Receiver   User    `json:"-"`
	ServiceID  uint    `json:"service_id" gorm:"column:service_id"`
	Service    Service `json:"-"`
	Amount     float64 `gorm:"not null"`
}
