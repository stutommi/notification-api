package main

import "os"

func main() {
	es := NewEmailService(
		emailConfig{
			sendGridApiKey: os.Getenv("SENDGRID_API_KEY"),
		})
	ss := &smsService{}
	ns := notificationService{emailService: es, smsService: ss}

	// Test email
	ed := EmailDetails{
		message:      "<h1>test from notification API</h1>",
		emailAddress: "<testemail@example.com>",
		html:         true,
	}

	ns.sendEmail(ed)
}
