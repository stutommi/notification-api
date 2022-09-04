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
	sendGridApiKey    string
	emailSenderDomain string
}

type emailService struct {
	cfg emailConfig
}

// Send email
func (es emailService) Send(ed EmailDetails) error {
	message, msgCreationError := es.createEmail(ed)
	if msgCreationError != nil {
		return fmt.Errorf("error creating email: %s", msgCreationError.Error())
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

// Validate email address
func (es emailService) validateEmail(emailAddress string) (*mail.Email, error) {
	email, err := mail.ParseEmail(emailAddress)
	return email, err
}

// Create email message
func (es emailService) createEmail(ed EmailDetails) (*mail.SGMailV3, error) {
	email, adressErr := es.validateEmail(ed.emailAddress)
	if adressErr != nil {
		return nil, fmt.Errorf("invalid email address: %s", adressErr.Error())
	}

	from := mail.NewEmail(ed.from, es.cfg.emailSenderDomain)
	to := mail.NewEmail(email.Name, email.Address)

	var emailMessage *mail.SGMailV3
	if ed.html {
		emailMessage = mail.NewSingleEmail(from, ed.subject, to, "", ed.message)
	} else {
		emailMessage = mail.NewSingleEmail(from, ed.subject, to, ed.message, "")
	}

	return emailMessage, nil
}

// initialise and return email service
func NewEmailService(cfg emailConfig) *emailService {
	es := emailService{
		cfg: cfg,
	}
	return &es
}
