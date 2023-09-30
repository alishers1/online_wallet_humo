package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Number  string  `json:"number" gorm:"unique;not null"`
	Balance float64 `json:"balance" gorm:"column:balance"`
	UserID  uint    `json:"user_id" gorm:"column:user_id"`
	User    User    `json:"-"`
}
