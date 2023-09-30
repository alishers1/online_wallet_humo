package configs

import (
	"encoding/json"
	"online_wallet_humo/pkg/models"
	"os"
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
