package config

import (
	"encoding/json"
	"log"
	"os"
)

const configFileName = ".gophzm.json"

type Config struct {
	DbURL string `json:"db_url"`
}

func Read() Config {
	filePath, err := getConfigFilePath()

	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	config := Config{}

	err = json.Unmarshal(jsonData, &config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

func write(cfg Config) error {
	jsonData, err := json.MarshalIndent(cfg, "", "  ")

	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(filePath)

	if err != nil {
		return nil
	}
	defer file.Close()

	_, err = file.Write(jsonData)

	if err != nil {
		return err
	}

	file.Sync()

	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	return homeDir + "/" + configFileName, nil
}
