package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetInputVars() ([]string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return []string{}, errors.New("error loading .env file")
	}

	return []string{
		os.Getenv("term"), os.Getenv("year"), os.Getenv("courseCode"),
	}, nil
}

func GetSenderEmail() ([]string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return []string{}, errors.New("error loading .env file")
	}

	return []string{
		os.Getenv("senderEmail"), os.Getenv("appPassword"),
	}, nil
}
