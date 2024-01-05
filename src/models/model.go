package models

type Section struct {
	CourseID              string `json:"courseId"`
	CourseName            string
	ClassSection          int    `json:"classSection"`
	CourseComponent       string `json:"courseComponent"`
	MaxEnrollmentCapacity int    `json:"maxEnrollmentCapacity"`
	EnrolledStudents      int    `json:"enrolledStudents"`
	HasSpace              bool
}
