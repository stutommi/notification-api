package main

import (
	"testing"
)

func TestInvalidEmail(t *testing.T) {
	es := &emailService{}
	invalidEmailAddress := "invalidaddress"

	_, err := es.validateEmail(invalidEmailAddress)

	if err == nil {
		t.Errorf("Expected error, got none")
	}
}

func TestValidEmail(t *testing.T) {
	es := &emailService{}
	validEmailAddress := "bob@example.com"

	_, err := es.validateEmail(validEmailAddress)

	if err != nil {
		t.Errorf("Expected error, got none")
	}
}
