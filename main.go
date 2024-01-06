package main

import (
	"fmt"
	"log"
	"time"

	"github.com/danjelhuang/course-notifier/src/models"
	"github.com/danjelhuang/course-notifier/src/network"
	"github.com/danjelhuang/course-notifier/src/sender"
	"github.com/danjelhuang/course-notifier/src/utils"
)

func worker(courses []models.Course) {
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

func main() {
	courses, err := utils.GetCourses()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(10 * time.Minute)
	for range ticker.C {
		worker(courses)
	}
}

// infinite loop, don't spam emails