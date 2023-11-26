package configs

import (
	"encoding/json"
	"fmt"
	"online_wallet_humo/pkg/models"
	"os"

	"github.com/sirupsen/logrus"
)

func InitConfigs() (*models.Config, error) {
	var configs models.Config

	bytes, err := os.ReadFile("../../configs/config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &configs)
	if err != nil {
		return nil, err
	}

	return &configs, nil
}

func GetConfigs() (serverStr, dbStr string, err error) {
	config, err := InitConfigs()
	if err != nil {
		logrus.Error(err)
		return "", "", err
	}

	address := config.Server.Host + config.Server.Port

	dbAddress := ToStringDBConfig(config)

	return address, dbAddress, nil
}

func ToStringDBConfig(c *models.Config) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Database.Host, c.Database.User, c.Database.Password, c.Database.DBName, c.Database.Port, c.Database.SSLMode,
	)
}
