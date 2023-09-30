package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func initDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=humo port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}

	return db
}

func StartDBConn() {
	database = initDB()
}

func GetDBConn() *gorm.DB {
	return database
}

func CloseDBConn() {
	sqlDb, err := GetDBConn().DB()
	if err != nil {
		logrus.Error(err)
	}

	sqlDb.Close()
}
