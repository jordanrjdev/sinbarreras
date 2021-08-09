package config

import (
	"encoding/json"
	"os"

	"github.com/JGurus/sin-barreras-backend/models"
)

func GetConfigDB() (models.ConfigDB, error) {
	config := models.ConfigDB{}
	file, err := os.Open("./config/database.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func GetConfigServer() (models.ConfigServer, error) {
	config := models.ConfigServer{}
	file, err := os.Open("./config/server.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
