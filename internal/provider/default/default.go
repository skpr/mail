package local

import (
	"context"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"golang.org/x/sync/errgroup"

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
	// // FallbackSMTPPort is the port used to connect to the SMTP server.
	// FallbackSMTPPort = "1025"
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

	// smtpPort := os.Getenv("SKPRMAIL_SMTP_PORT")
	// if smtpPort == "" {
	// 	smtpPort = FallbackSMTPPort
	// }

	// tlsConfig := &tls.Config{
	// 	InsecureSkipVerify: true,
	// 	ServerName:         "",
	// }

	client, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to dial smtp server %w", err)
	}

	// Connect to the SMTP server
	// @todo, Should reuse addr and make the port configurable.
	// client, err := smtp.NewClient(conn, addr)
	// if err != nil {
	// 	return fmt.Errorf("failed to dial mailhog smtp %w", err)
	// }

	// if err = client.Auth(auth); err != nil {
	// 	log.Panic(err)
	// }

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	eg := errgroup.Group{}

	// Sending the email.
	eg.Go(func() error {
		defer cancel()

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

		_, err = wr.Write([]byte(data))
		if err != nil {
			log.Panic(err)
		}

		// @todo, Do the sending.

		return nil
	})

	// Closing the client.
	eg.Go(func() error {
		<-ctx.Done()

		err := client.Close()
		if err != nil {
			return fmt.Errorf("failed to close: %w", err)
		}

		// err = client.Quit()
		// if err != nil {
		// 	return fmt.Errorf("failed to quit: %w", err)
		// }

		return nil
	})

	log.Println("successfully sent message via mailhog smtp")

	return eg.Wait()
}
