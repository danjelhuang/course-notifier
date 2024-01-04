package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/danjelhuang/course-notifier/src/network"
)

func readInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter %s: ", prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func parseInput() (string, string, string, error) {
	term, err := readInput("term")
	if err != nil {
		return "", "", "", err
	}

	year, err := readInput("year")
	if err != nil {
		return "", "", "", err
	}

	courseCode, err := readInput("course code")
	if err != nil {
		return "", "", "", err
	}
	return term, year, courseCode, nil
}

func main() {
	term, year, courseCode, err := parseInput()
	if err != nil {
		log.Fatal(err)
	}

	err = network.RequestAPI(term, year, courseCode)
	if err != nil {
		log.Fatal(err)
	}
}
