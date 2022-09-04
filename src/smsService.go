package main

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type smsConfig struct {
	twilioAccountSid  string
	twilioAuthToken   string
	twilioPhoneNumber string
}

type smsService struct {
	twilioClient      *twilio.RestClient
	twilioPhoneNumber string
}

// Send SMS
func (ss smsService) Send(sd SmsDetails) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo(sd.phoneNumber)
	params.SetFrom(ss.twilioPhoneNumber)
	params.SetBody(sd.message)

	_, err := ss.twilioClient.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("error sending sms: %s", err.Error())
	}

	return nil
}

// initialise and return email service
func NewSmsService(cfg smsConfig) *smsService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.twilioAccountSid,
		Password: cfg.twilioAuthToken,
	})

	ss := smsService{
		twilioClient:      client,
		twilioPhoneNumber: cfg.twilioPhoneNumber,
	}
	return &ss
}
