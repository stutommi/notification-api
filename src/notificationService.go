package main

import (
	"log"
)

type EmailDetails struct {
	emailAddress string
	message      string
	html         bool
}

// - - - Depedencies - - - //

type EmailSender interface {
	// Send single email message to specified email address(es).
	Send(emailDetails EmailDetails) error
}

type SmsSender interface {
	// Send single phone message to specified phone number.
	Send(phoneNumber string, message string) error
}

// - - - Interface - - - //

type Notificationer interface {
	// Send single email message to specified email address(es).
	sendEmail(emailDetails EmailDetails) error
	// Send single phone message to specified phone number.
	sendSms(phoneNumber string, message string) error
}

// - - - Implementation - - - //

type notificationService struct {
	emailService EmailSender
	smsService   SmsSender
}

func (ns *notificationService) sendSms(phoneNumber string, message string) error {
	err := ns.smsService.Send(phoneNumber, message)
	if err != nil {
		log.Println("sms sent succesfully")
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
