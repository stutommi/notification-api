package main

import "os"

func main() {
	es := NewEmailService(
		emailConfig{
			sendGridApiKey:    os.Getenv("SENDGRID_API_KEY"),
			emailSenderDomain: os.Getenv("EMAIL_SENDER_DOMAIN"),
		})

	ss := NewSmsService(smsConfig{
		twilioAccountSid:  os.Getenv("TWILIO_ACCOUNT_SID"),
		twilioAuthToken:   os.Getenv("TWILIO_AUTH_TOKEN"),
		twilioPhoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
	})
	ns := notificationService{emailService: es, smsService: ss}

	// Test email
	ed := EmailDetails{
		from:         "Elon Musk",
		subject:      "I want to give you my money",
		message:      "<h1>please take it</h1>",
		emailAddress: "<address>",
		html:         true,
	}

	// Test SMS
	sd := SmsDetails{
		message:     "greeting from notification-api",
		phoneNumber: "<phonenumber>",
	}

	ns.sendEmail(ed)
	ns.sendSms(sd)
}
