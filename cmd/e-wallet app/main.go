package main

import (
	"log"
	"online_wallet_humo/internal/db"
	"online_wallet_humo/internal/handler"
	"online_wallet_humo/logs"
	"online_wallet_humo/pkg/models"
	"online_wallet_humo/seeders"

	"github.com/sirupsen/logrus"
)

func main() {

	if err := logs.InitLogger(); err != nil {
		return
	}

	logrus.Error("This is the test")

	db.StartDBConn()

	err := db.GetDBConn().AutoMigrate(&models.User{}, &models.Service{}, &models.Card{}, &models.Transaction{}, &models.Favorite{})
	if err != nil {
		log.Println("Error while migration", err.Error())
		return
	}

	seeders.SeedUsers(db.GetDBConn())
	seeders.SeedServices(db.GetDBConn())
	seeders.SeedCards(db.GetDBConn())

	if err := handler.InitRoutes(); err != nil {
		log.Println(err)
		return
	}

}
