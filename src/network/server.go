package network

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/danjelhuang/course-notifier/src/models"
)

func setRequestBody(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")

	apiKey, err := getAPIKey("UW_API_KEY")
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", apiKey)

	return req, nil
}

func callAPI(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func RequestAPI(term, year, courseCode string) ([]byte, error) {
	termNumber, err := getTermNumber(term, year)
	if err != nil {
		return []byte{}, errors.New("term number error")
	}

	url := getURL(termNumber, courseCode)
	req, err := setRequestBody(url)
	if err != nil {
		return []byte{}, errors.New("set request body error")
	}

	body, err := callAPI(req)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func GetSections(body []byte) ([]models.Section, error) {
	var sections []models.Section
	err := json.Unmarshal(body, &sections)
	if err != nil {
		return []models.Section{}, errors.New("unmarshal error")
	}

	sections = filterSections(sections)
	return sections, nil
}
