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

	for _, course := range courses {
		sections, err := network.RequestAPI(course)
		if err != nil {
			log.Fatal(err)
		}

		for _, section := range sections {
			fmt.Println(section.CourseID, section.CourseName, section.ClassSection, section.CourseComponent, section.EnrolledStudents, section.MaxEnrollmentCapacity, section.HasSpace)

			if section.HasSpace {
				sender.SendEmail(section)
			}
		}
	}
}
