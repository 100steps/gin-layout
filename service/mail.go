package service

import (
	"strconv"

	"github.com/forseason/env"
	"github.com/go-gomail/gomail"
)

type MailService struct {
	serverHost   string
	serverPort   int
	fromEmail    string
	fromPassword string
	fromUsername string
}

func NewMailService() (*MailService, error) {
	port, err := strconv.Atoi(env.Get(("EMAIL_SERIVER_PORT"), ""))
	if err != nil {
		return nil, err
	}
	return &MailService{
		serverHost:   env.Get("EMAIL_SERVER_HOST", ""),
		serverPort:   port,
		fromEmail:    env.Get("EMAIL_FROM_EMAIL", ""),
		fromUsername: env.Get("EMAIL_FROM_USERNAME", ""),
		fromPassword: env.Get("EMAIL_FROM_PASSWORD", ""),
	}, nil
}

func (this *MailService) Send(body string, subject string, toers []string) error {
	if len(toers) == 0 {
		return nil
	}
	msg := gomail.NewMessage()
	msg.SetHeader("To", toers...)
	msg.SetHeader("From", this.fromUsername)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	dialer := gomail.NewDialer(this.serverHost, this.serverPort, this.fromUsername, this.fromPassword)
	return dialer.DialAndSend(msg)
}
