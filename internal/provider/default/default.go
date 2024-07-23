package local

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"time"

	"github.com/skpr/mail/internal/mailutils"
)

const (
	// EnvAddr used to configure the address where mail is sent.
	EnvAddr = "SKPRMAIL_ADDR"
	// EnvFrom used to configure the FROM address appled to mail.
	EnvFrom = "SKPRMAIL_FROM"
	// FallbackAddr where mail will be forwarded to.
	FallbackAddr = "localhost:1025"
	// FallbackFrom address which will be applied to email.
	FallbackFrom = "skprmail@skpprmail.com"
)

// Send the email to Mailhog.
func Send(ctx context.Context, to []string, msg *mail.Message) error {
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

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Printf("sleeping for 1 second")
			select {
			case <-ctx.Done():
				fmt.Printf("cancelling")
				os.Exit(1)
			}
		}
	}()

	fmt.Printf("dialing smtp server %s", addr)
	client, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to dial smtp server %w", err)
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("failed to add from address %w", err)
	}

	for _, recipient := range to {
		err = client.Rcpt(recipient)
		if err != nil {
			return fmt.Errorf("failed to add recipient: %w", err)
		}
	}

	wr, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to initiate writer: %w", err)
	}

	log.Println("writing message to mailhog smtp")
	_, err = wr.Write([]byte(data))
	if err != nil {
		log.Panic(err)
	}
	log.Println("successfully wrote message to mailhog smtp")
	err = client.Close()
	if err != nil {
		return fmt.Errorf("failed to close: %w", err)
	}
	log.Println("successfully sent message via mailhog smtp")
	// time.Sleep(20 * time.Second)

	return nil
}
