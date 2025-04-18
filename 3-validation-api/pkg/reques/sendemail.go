package reques

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/storage"

	"net/smtp"

	"github.com/jordan-wright/email"
)

type SenderMail struct {
	TO string
}

func MailSend(emailLoad *storage.EmailList, sender *configs.Config, bodyMail string) error {

	e := email.NewEmail()
	e.From = sender.Email
	e.To = []string{emailLoad.Mail}

	e.Subject = "Awesome Subject"
	e.Text = []byte(bodyMail)

	err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", sender.Email, sender.Password, "smtp.gmail.com"))
	if err != nil {
		return err
	}
	return nil
}
