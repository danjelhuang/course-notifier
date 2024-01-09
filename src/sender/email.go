package sender

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/danjelhuang/course-notifier/src/models"
	"github.com/danjelhuang/course-notifier/src/utils"
)

func SendEmail(section models.Section, receivers []string) {
	senderEmail, err := utils.GetSenderEmail()
	if err != nil {
		log.Fatal(err)
	}
	from := senderEmail[0]
	password := senderEmail[1]
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	courseName := strings.ToUpper(section.CourseName)

	for _, receiver := range receivers {
		message := fmt.Sprintf("Subject: %s Opening\r\nTo: %s\r\n\r\n %s has an opening in Section %d", courseName, receiver, courseName, section.ClassSection)

		auth := smtp.PlainAuth("", from, password, smtpHost)

		err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{receiver}, []byte(message))
		if err != nil {
			log.Fatal(err)
		}
	}
}
