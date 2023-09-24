package common

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"os"
)

type MailConfig struct {
	Username   string
	Password   string
	Host       string
	Port       int
	To         string
	Subject    string
	AttachPath string
	Body       string
}

func SendMail(email MailConfig) {
	m := gomail.NewMessage()
	m.SetHeader("From", email.Username)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Body)
	m.Attach(email.AttachPath)
	d := gomail.NewDialer(email.Host, email.Port, email.Username, email.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: d.Host}
	if err := d.DialAndSend(m); err != nil {

	} else {
		os.Remove(email.AttachPath)
	}
}
