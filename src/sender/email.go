package sender

import (
	"log"
	"net/smtp"
	"strconv"

	"github.com/danjelhuang/course-notifier/src/models"
	"github.com/danjelhuang/course-notifier/src/utils"
)

func SendEmail(section models.Section) {
	senderEmail, err := utils.GetSenderEmail()
	if err != nil {
		log.Fatal(err)
	}
	from := senderEmail[0]
	password := senderEmail[1]
	to := "danielhuang18@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := "Subject: Course Opening\r\n" +
		"To: " + to + "\r\n" +
		"\r\n" +
		"Section " + strconv.Itoa(section.ClassSection) + " of " + " now has space."

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
