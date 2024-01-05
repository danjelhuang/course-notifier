package network

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/danjelhuang/course-notifier/src/models"
	"github.com/joho/godotenv"
)

func getTermStartMonth(term string) (int, error) {
	startMonths := map[string]int{"winter": 1, "spring": 5, "summer": 5, "fall": 9}

	startMonthNumber, ok := startMonths[term]
	if ok {
		return startMonthNumber, nil
	} else {
		return 0, errors.New("invalid term")
	}
}

func getTermNumber(term, year string) (string, error) {
	termStartMonth, err := getTermStartMonth(term)
	if err != nil {
		return "", errors.New("invalid term")
	}

	yearNumber, err := strconv.Atoi(year)
	if err != nil {
		return "", errors.New("invalid year")
	}

	var termNumber string
	if yearNumber < 2000 {
		termNumber = fmt.Sprintf("0%s%d", year[len(year)-2:], termStartMonth)
	} else {
		termNumber = fmt.Sprintf("1%s%d", year[len(year)-2:], termStartMonth)
	}
	return termNumber, nil
}

func getURL(termNumber, courseCode string) string {
	const baseURL = "https://openapi.data.uwaterloo.ca/v3/ClassSchedules"
	courseSubject := strings.Split(courseCode, " ")[0]
	courseNumber := strings.Split(courseCode, " ")[1]

	url := fmt.Sprintf("%s/%s/%s/%s", baseURL, termNumber, courseSubject, courseNumber)
	return url
}

func getAPIKey(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("error loading .env file")
	}

	return os.Getenv(key), nil
}

func filterSections(sections []models.Section) []models.Section {
	var filtered []models.Section
	const lecture = "LEC"
	for _, section := range sections {
		if section.CourseComponent == lecture {
			filtered = append(filtered, section)
		}
	}
	return filtered
}
