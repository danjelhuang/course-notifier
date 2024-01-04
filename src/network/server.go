package network

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getAPIKey(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("error loading .env file")
	}

	return os.Getenv(key), nil
}

func RequestAPI(term, year, courseCode string) error {
	client := &http.Client{}

	termNumber, err := getTermNumber(term, year)
	if err != nil {
		return errors.New("term number error")
	}

	url := getURL(termNumber, courseCode)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "application/json")

	apiKey, err := getAPIKey("UW_API_KEY")
	if err != nil {
		return err
	}
	req.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", bodyText)
	return nil
}
