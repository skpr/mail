package mailhog

import (
	"bytes"
	"net/mail"
	"net/smtp"
)

const (
	// Addr which Mailhog will receive mail.
	Addr = "mailhog:1025"
	// From address which will be applied to email.
	From = "skprmail"
)

// Send the email to Mailhog.
func Send(data []byte) error {
	msg, err := mail.ReadMessage(bytes.NewReader(data))
	if err != nil {
		return err
	}

	to := msg.Header.Get("To")

	return smtp.SendMail(Addr, nil, From, []string{to}, data)
}
