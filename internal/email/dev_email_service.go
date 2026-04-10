package email

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

type DevEmailService struct {
	from   string
	writer io.Writer
}

func NewDevEmailService(from string, writer io.Writer) *DevEmailService {
	return &DevEmailService{
		from:   from,
		writer: writer,
	}
}

func (s *DevEmailService) Send(ctx context.Context, email Email) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	var buf bytes.Buffer

	fmt.Fprintf(&buf, "From: %s\n", s.from)
	fmt.Fprintf(&buf, "To: %s\n", email.To)
	fmt.Fprintf(&buf, "Subject: %s\n\n", email.Subject)
	fmt.Fprintf(&buf, "Body: %s\n", email.Body)

	if _, err := s.writer.Write(buf.Bytes()); err != nil {
		return ErrEmailWriteError
	}

	return nil
}
