package mocking

import "testing"

func TestSendWelcomeEmail(t *testing.T) {
	mock := &MockSender{}

	SendWelcomeEmail(mock, "user@gmail.com")

	if len(mock.SentEmails) != 1 {
		t.Errorf("expected 1 email, got: %v", len(mock.SentEmails))
	}

	if mock.SentEmails[0] != "user@gmail.com" {
		t.Errorf("expected user@gmail.com, got %s", mock.SentEmails[0])
	}
}
