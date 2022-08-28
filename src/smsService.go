package main

type smsService struct {
}

func (ss smsService) Send(phoneNumber string, message string) error {
	return nil
}
