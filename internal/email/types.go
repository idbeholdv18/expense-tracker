package email

import "context"

type Email struct {
	To      string
	Subject string
	Body    []byte
}

type EmailService interface {
	Send(ctx context.Context, email Email) error
}
