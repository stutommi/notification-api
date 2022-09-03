package main

import (
	"encoding/json"
	"fmt"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendMailResponse struct {
	Errors []struct {
		Message string      `json:"message"`
		Field   string      `json:"field"`
		Help    interface{} `json:"help"`
	} `json:"errors"`
}

type emailConfig struct {
	sendGridApiKey string
}

type emailService struct {
	cfg emailConfig
}

func (es emailService) Send(ed EmailDetails) error {
	email, adressErr := es.validateEmail(ed.emailAddress)
	if adressErr != nil {
		return fmt.Errorf("invalid email address: %s", adressErr.Error())
	}

	from := mail.NewEmail("musignalis", "tommi.teetee@hotmail.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail(email.Name, email.Address)

	var message *mail.SGMailV3
	if ed.html {
		message = mail.NewSingleEmail(from, subject, to, "", ed.message)
	} else {
		message = mail.NewSingleEmail(from, subject, to, ed.message, "")
	}

	client := sendgrid.NewSendClient(es.cfg.sendGridApiKey)
	response, err := client.Send(message)

	if err != nil {
		return fmt.Errorf("error sending email: %s", err.Error())
	} else if response.StatusCode != 202 {
		var result SendMailResponse
		json.Unmarshal([]byte(response.Body), &result)
		return fmt.Errorf("failed sending email: %s", result.Errors[0].Message)
	}

	return nil
}

func (es emailService) validateEmail(emailAddress string) (*mail.Email, error) {
	email, err := mail.ParseEmail(emailAddress)
	return email, err
}

func NewEmailService(cfg emailConfig) *emailService {
	es := emailService{
		cfg: cfg,
	}
	return &es
}
