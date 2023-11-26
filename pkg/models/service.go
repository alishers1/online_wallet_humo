package models

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	IsActive bool   `gorm:"default:true"`
	Name     string `json:"name" gorm:"unique"`
}
