package config

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/zgierz/klimson/backend/models"
)

func GetConfig(mainPath string, serverConfigPath string) (*models.Config, error) {
	file, err := os.Open(mainPath + serverConfigPath + "/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cfg = PreGeneratedConfig
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}

func GetSpecificConfigRecord() {

}
