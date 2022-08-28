package main

import "os"

func main() {
	es := NewEmailService(
		emailConfig{
			from: os.Getenv("smtp_from"),
			pwd:  os.Getenv("smtp_pwd"),
			host: os.Getenv("smtp_host"),
			port: os.Getenv("smtp_port"),
		})
	ss := &smsService{}
	println(os.Getenv("smtp_from"))
	ns := notificationService{emailService: es, smsService: ss}

	// Test email
	ed := EmailDetails{
		message:      "test from notification API",
		emailAddress: []string{"tommi.teetee@hotmail.com"},
	}

	ns.sendEmail(ed)
}
