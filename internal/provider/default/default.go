package local

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/skpr/mail/internal/mailutils"
)

const (
	// EnvAddr used to configure the address where mail is sent.
	EnvAddr = "SKPR_MAIL_ADDR"
	// EnvFrom used to configure the FROM address appled to mail.
	EnvFrom = "SKPR_MAIL_FROM"
	// FallbackAddr where mail will be forwarded to.
	FallbackAddr = "mail:1025"
	// FallbackFrom address which will be applied to email.
	FallbackFrom = "skprmail"
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

	addr := os.Getenv(EnvAddr)
	if addr == "" {
		addr = FallbackAddr
	}

	from := os.Getenv(EnvFrom)
	if addr == "" {
		addr = FallbackFrom
	}

	err = smtp.SendMail(addr, nil, from, to, data)
	if err != nil {
		return fmt.Errorf("failed to send message via mailhog smtp %w", err)
	}

	log.Println("successfully sent message via mailhog smtp")

	return nil
}
