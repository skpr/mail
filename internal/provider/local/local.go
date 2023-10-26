package local

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"

	"github.com/skpr/mail/internal/mailutils"
)

const (
	// From address which will be applied to email.
	From = "skprmail"
)

// Send the email to local mail server.
func Send(addr string, to []string, msg *mail.Message) error {
	data, err := mailutils.MessageToBytes(msg)
	if err != nil {
		return err
	}
	if val, ok := msg.Header[mailutils.HeaderTo]; ok {
		to = append(to, val...)
	}

	err = smtp.SendMail(addr, nil, From, to, data)
	if err != nil {
		return fmt.Errorf("failed to send message via smtp %w", err)
	}
	log.Println("successfully sent message via smtp")

	return nil
}
