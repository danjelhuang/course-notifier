package main

import (
	"fmt"
	"log"

	"github.com/danjelhuang/course-notifier/src/network"
	"github.com/danjelhuang/course-notifier/src/utils"
)

func main() {
	inputs, err := utils.GetInputVars()
	if err != nil {
		log.Fatal(err)
	}

	body, err := network.RequestAPI(inputs)
	if err != nil {
		log.Fatal(err)
	}

	sections, err := network.GetSections(body)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range sections {
		fmt.Println(p.CourseID, p.ClassSection, p.CourseComponent, p.EnrolledStudents, p.MaxEnrollmentCapacity)
	}
}

// TODO:
//		 finish logic off those structs
// 		 send notifications when something happens
