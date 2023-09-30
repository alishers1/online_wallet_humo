package seeders

import (
	"log"
	"online_wallet_humo/pkg/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedCards(db *gorm.DB) {
	cards := []models.Card{
		{
			Number: "5058270200001111",
			Balance: 4000,
		},
		{
			Number: "5058270211112222",
			Balance: 2400,
		},
	}

	for _, card := range cards {
		if err := db.Create(&card).Error; err != nil {
			return
		}
	}
}

func SeedServices(db *gorm.DB) {
	services := []models.Service{
		{
			Name: "From card to wallet",
		},
		{
			Name: "From wallet to wallet",
		},
		{
			Name: "Universal service",
		},
	}

	for _, service := range services {
		if err := db.Create(&service).Error; err != nil {
			return
		}
	}
}

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{
			Balance:      3000,
			FullName:     "Siddiqov Toji",
			PhoneNumber:  "907596938",
			Email:        "test@gmail.com",
			PasswordHash: string(generateFromPassword("54321")),
		},
		{
			Balance:      2000,
			FullName:     "Siddiqov Aslam",
			PhoneNumber:  "939392163",
			Email:        "test1@gmail.com",
			PasswordHash: string(generateFromPassword("12345")),
		},
		{
			Balance:      200,
			FullName:     "Siddiqov Alisher",
			PhoneNumber:  "939392166",
			Email:        "test2@gmail.com",
			PasswordHash: string(generateFromPassword("12345")),
			IsAdmin:      true,
		},
		{
			Balance:      300,
			FullName:     "Siddiqov Atorud",
			PhoneNumber:  "939392164",
			Email:        "test3@gmail.com",
			PasswordHash: string(generateFromPassword("12345")),
		},
		{
			Balance:      400,
			FullName:     "Siddiqov Ashraf",
			PhoneNumber:  "939392160",
			Email:        "test3@gmail.com",
			PasswordHash: string(generateFromPassword("12345")),
		},
		{
			Balance:      500,
			FullName:     "Fozilov Rustam",
			PhoneNumber:  "939392161",
			Email:        "test2@gmail.com",
			PasswordHash: string(generateFromPassword("12345")),
		},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return
		}
	}
}

func generateFromPassword(password string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return hash
}
