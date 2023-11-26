package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Balance      float64 `json:"balance" gorm:"column:balance"`
	FullName     string  `json:"full_name" gorm:"column:full_name"`
	PhoneNumber  string  `json:"phone_number" gorm:"unique"`
	Email        string  `json:"email" gorm:"type:varchar(100);unique_index"`
	PasswordHash string  `json:"password_hash" gorm:"column:password_hash"`
	Avatar       string  `json:"avatar" gorm:"column:avatar"`
	IsActive     bool    `json:"is_active" gorm:"default:true"`
	IsAdmin      bool    `json:"is_admin" gorm:"default:false"`
}
