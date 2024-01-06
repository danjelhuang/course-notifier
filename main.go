package main

import (
	"fmt"
	"log"

	"github.com/danjelhuang/course-notifier/src/network"
	"github.com/danjelhuang/course-notifier/src/sender"
	"github.com/danjelhuang/course-notifier/src/utils"
)

func main() {
	courses, err := utils.GetCourses()
	if err != nil {
		log.Fatal(err)
	}

	sections, err := network.RequestAPI(courses)
	if err != nil {
		log.Fatal(err)
	}

	for _, section := range sections {
		fmt.Println(section.CourseID, section.ClassSection, section.CourseComponent, section.EnrolledStudents, section.MaxEnrollmentCapacity, section.HasSpace)

		if section.HasSpace {
			sender.SendEmail(section)
		}
	}
}
