package service

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"path"

	"github.com/Cmdliner/streem/internal/config"
	gomail "gopkg.in/mail.v2"
)

type EmailService struct {
	cfg  *config.Config
}

func NewEmailService(cfg *config.Config) *EmailService {
	return &EmailService{
		cfg,
	}
}

func (r EmailService) SendEmail(sender, recepient, subject string, cfg *config.Config) {

	_, err := template.ParseFiles(path.Clean("../static/index.html"))

	message := gomail.NewMessage()

	message.SetHeader("From", sender)
	message.SetHeader("To", recepient)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", `<html>This is a test email</html>`)

	if err != nil {
		panic(err)
	}

	// Set up the SMTP dialer
	dialer := gomail.NewDialer(cfg.Email.Provider, 587, cfg.Email.Username, cfg.Email.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error: %w", err)
		panic(err)
	} else {
		fmt.Println("Email sent successfully")
	}
}
