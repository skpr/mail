package local

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	skprconfig "github.com/skpr/go-config"
	"github.com/skpr/mail/internal/mailutils"
)

const (
	// EnvAddr used to configure the address where mail is sent.
	EnvAddr = "SKPRMAIL_ADDR"
	// EnvFrom used to configure the FROM address appled to mail.
	EnvFrom = "SKPRMAIL_FROM"
	// EnvPort used to configure the port for the destination address.
	EnvPort = "SKPRMAIL_PORT"
	// FallbackAddr where mail will be forwarded to.
	FallbackAddr = "mail"
	// FallbackFrom address which will be applied to email.
	FallbackFrom = "skprmail"
	// FallbackPort address which will be applied to email.
	FallbackPort = "1025"
	// ConfigAddr used to override and set the address where mail is sent using a Skpr config.
	ConfigAddr = "smtp.hostname"
	// ConfigFrom used to override and set the FROM address where mail is sent using a Skpr config.
	ConfigFrom = "smtp.from.address"
	// ConfigPort used to override and set the address port where mail is sent using a Skpr config.
	ConfigPort = "smtp.port"
)

// Send the email to Mailhog.
func Send(ctx context.Context, to []string, msg *mail.Message) error {
	// The GO SMTP package is difficult to cancel using context.
	// This provider should only ever be used for local development tasks.
	go func() {
		<-ctx.Done()
		fmt.Println("Context cancelled")
		os.Exit(1)
	}()

	config, err := skprconfig.Load()
	if err != nil && !errors.Is(err, skprconfig.ErrNotFound) {
		panic(err)
	}

	data, err := mailutils.MessageToBytes(msg)
	if err != nil {
		return err
	}

	if val, ok := msg.Header[mailutils.HeaderTo]; ok {
		to = append(to, val...)
	}

	addr := os.Getenv(EnvAddr)
	if addr == "" {
		addr = config.GetWithFallback(ConfigAddr, FallbackAddr)
	}

	port := os.Getenv(EnvPort)
	if port == "" {
		port = config.GetWithFallback(ConfigPort, FallbackPort)
	}

	from := os.Getenv(EnvFrom)
	if from == "" {
		from = config.GetWithFallback(ConfigFrom, FallbackFrom)
	}

	destination := fmt.Sprintf("%s:%s", addr, port)

	err = smtp.SendMail(destination, nil, from, to, data)
	if err != nil {
		return fmt.Errorf("failed to send message via mailhog smtp %w", err)
	}

	log.Println("successfully sent message via mailhog smtp")

	return nil
}
