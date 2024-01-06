package utils

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/danjelhuang/course-notifier/src/models"
	"github.com/joho/godotenv"
)

const configFile = "config.json"

func loadConfig() (models.Config, error) {
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return models.Config{}, err
	}

	var config models.Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return models.Config{}, errors.New("unmarshal error")
	}
	return config, nil
}

func GetCourses() ([]models.Course, error) {
	config, err := loadConfig()
	if err != nil {
		return []models.Course{}, errors.New("load config error")
	}

	return config.Courses, nil
}

func GetSenderEmail() ([]string, error) {
	config, err := loadConfig()
	if err != nil {
		return []string{}, errors.New("load config error")
	}

	err = godotenv.Load(".env")
	if err != nil {
		return []string{}, errors.New("error loading .env file")
	}

	return []string{
		config.SenderEmail, os.Getenv("appPassword"),
	}, nil
}
