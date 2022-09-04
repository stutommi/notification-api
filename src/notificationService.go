package main

import (
	"log"
)

type EmailDetails struct {
	emailAddress string
	subject      string
	from         string
	message      string
	html         bool
}

type SmsDetails struct {
	phoneNumber string
	message     string
}

// - - - Depedencies - - - //

type EmailSender interface {
	// Send single email message to specified email address(es).
	Send(emailDetails EmailDetails) error
}

type SmsSender interface {
	// Send single phone message to specified phone number.
	Send(smsDetails SmsDetails) error
}

// - - - Interface - - - //

type Notificationer interface {
	// Send single email message to specified email address(es).
	sendEmail(emailDetails EmailDetails) error
	// Send single phone message to specified phone number.
	sendSms(smsDetails SmsDetails) error
}

// - - - Implementation - - - //

type notificationService struct {
	emailService EmailSender
	smsService   SmsSender
}

func (ns *notificationService) sendSms(sd SmsDetails) error {
	err := ns.smsService.Send(sd)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("sms sent succesfully")
	return nil
}

func (ns *notificationService) sendEmail(emailDetails EmailDetails) error {
	err := ns.emailService.Send(emailDetails)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("email sent succesfully")
	return nil
}
