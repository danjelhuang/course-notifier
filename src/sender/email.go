package sender

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

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
	courseName := strings.ToUpper(section.CourseName)

	message := fmt.Sprintf("Subject: %s Opening\r\nTo: %s\r\n\r\n %s has an opening in Section %d", courseName, to, courseName, section.ClassSection)

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
