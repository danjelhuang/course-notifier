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

var sentEmails = map[string]bool{}

func worker(courses []models.Course) {
	for _, course := range courses {
		sections, err := network.RequestAPI(course)
		if err != nil {
			log.Fatal(err)
		}

		for _, section := range sections {
			fmt.Println(section.CourseID, section.CourseName, section.ClassSection, section.CourseComponent, section.EnrolledStudents, section.MaxEnrollmentCapacity, section.HasSpace)
			sectionID := fmt.Sprintf("%s %d", section.CourseName, section.ClassSection)

			if _, sent := sentEmails[sectionID]; section.HasSpace && !sent {
				sender.SendEmail(section, course.ReceiverEmails)
				sentEmails[sectionID] = true
			}
		}
	}
}

func main() {
	courses, err := utils.GetCourses()
	if err != nil {
		log.Fatal(err)
	}

	interval, err := utils.GetInterval()
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	for range ticker.C {
		worker(courses)
	}
}

// when to remove from sentEmails
