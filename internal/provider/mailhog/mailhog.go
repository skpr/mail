package mailhog

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/skpr/mail/internal/mailutils"
)

const (
	// Addr which Mailhog will receive mail.
	Addr = "mailhog:1025"
	// From address which will be applied to email.
	From = "skprmail"
)

// Send the email to Mailhog.
func Send(to []string, msg *mail.Message) error {
	data, err := mailutils.MessageToBytes(msg)
	if err != nil {
		return err
	}
	if val, ok := msg.Header[mailutils.HeaderTo]; ok {
		to = append(to, val...)
	}

	err = smtp.SendMail(Addr, nil, From, to, data)
	if err != nil {
		return fmt.Errorf("failed to send message via mailhog smtp %w", err)
	}
	log.Println("successfully sent message via mailhog smtp")

	return nil
}
