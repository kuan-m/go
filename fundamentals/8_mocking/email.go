package mocking

// 1. Interface
type EmailSender interface {
	SendEmail(to, body string) error
}

// Struct который имплементит interface
type RealSender struct {
}

func (s *RealSender) SendEmail(to, body string) error {
	// example: реальная отправка email чере зsmtp пакет
	return nil
}

// Mocker (Притворщик) имплементит interface
type MockSender struct {
	SentEmails []string
}

func (s *MockSender) SendEmail(to, body string) error {
	s.SentEmails = append(s.SentEmails, to)
	return nil
}

// 2. Функция с DI
func SendWelcomeEmail(sender EmailSender, userEmail string) {
	sender.SendEmail(userEmail, "Welcome Hello!")
}
