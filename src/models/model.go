package models

type Config struct {
	SenderEmail string   `json:"sender_email"`
	Courses     []Course `json:"courses"`
}

type Course struct {
	Term       string `json:"term"`
	Year       string `json:"year"`
	CourseCode string `json:"course_code"`
}

type Section struct {
	CourseID              string `json:"courseId"`
	CourseName            string
	ClassSection          int    `json:"classSection"`
	CourseComponent       string `json:"courseComponent"`
	MaxEnrollmentCapacity int    `json:"maxEnrollmentCapacity"`
	EnrolledStudents      int    `json:"enrolledStudents"`
	HasSpace              bool
}
