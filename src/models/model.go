package models

type Config struct {
	SenderEmail     string   `json:"sender_email"`
	IntervalMinutes int      `json:"interval_minutes"`
	Courses         []Course `json:"courses"`
}

type Course struct {
	Term           string   `json:"term"`
	Year           string   `json:"year"`
	CourseCode     string   `json:"course_code"`
	ReceiverEmails []string `json:"receiver_emails"`
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
