package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var (
		term       string
		year       string
		courseCode string
	)

	fmt.Println("Enter term: ")
	term, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	term = strings.TrimSpace(term)

	fmt.Println("Enter year: ")
	year, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	year = strings.TrimSpace(year)

	fmt.Println("Enter course code: ")
	courseCode, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	courseCode = strings.TrimSpace(courseCode)

	fmt.Println(term)
	fmt.Println(year)
	fmt.Println(courseCode)
}
