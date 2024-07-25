package local

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/skpr/mail/internal/mailutils"
)

const (
	// EnvAddr used to configure the address where mail is sent.
	EnvAddr = "SKPRMAIL_ADDR"
	// EnvFrom used to configure the FROM address appled to mail.
	EnvFrom = "SKPRMAIL_FROM"
	// FallbackAddr where mail will be forwarded to.
	FallbackAddr = "mail:1025"
	// FallbackFrom address which will be applied to email.
	FallbackFrom = "mail@skpr.localhost"
)

// Send the email to Mailhog.
func Send(ctx context.Context, to []string, msg *mail.Message) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// The GO SMTP package is difficult to cancel using context.
	// This provider should only ever be used for local development tasks.
	go func() {
		<-ctx.Done()
		fmt.Printf("Contacting the local smtp server timed out, cancelling...\n")
		os.Exit(1)
	}()

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
	if from == "" {
		from = FallbackFrom
	}

	err = smtp.SendMail(addr, nil, from, to, data)
	if err != nil {
		return fmt.Errorf("failed to send message via mailhog smtp %w", err)
	}
	log.Println("successfully sent message via mailhog smtp")
	return nil
}
