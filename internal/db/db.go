package db

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func initDB(dbStr string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dbStr), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}

	return db
}

func StartDBConn(dbStr string) {
	database = initDB(dbStr)
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
