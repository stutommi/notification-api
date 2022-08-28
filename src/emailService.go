package main

import (
	"fmt"
	"net/mail"
	"net/smtp"
)

type emailConfig struct {
	from string
	host string
	port string
	pwd  string
}

type emailService struct {
	cfg  emailConfig
	auth smtp.Auth
}

func (es emailService) Send(ed EmailDetails) error {
	adressErr := es.validateEmail(ed.emailAddress[0])
	if adressErr != nil {
		return fmt.Errorf("invalid email address: %s", adressErr.Error())
	}

	messageBytes := []byte(ed.message)
	err := smtp.SendMail(es.cfg.host+":"+es.cfg.port, es.auth, es.cfg.from, ed.emailAddress, messageBytes)
	if err != nil {
		return fmt.Errorf("error sending email: %s", err.Error())
	}

	return nil
}

func (es emailService) validateEmail(emailAddress string) error {
	_, err := mail.ParseAddress(emailAddress)
	return err
}

func NewEmailService(cfg emailConfig) *emailService {
	auth := smtp.PlainAuth("", cfg.from, cfg.pwd, cfg.host)
	es := emailService{
		cfg:  cfg,
		auth: auth,
	}
	return &es
}
