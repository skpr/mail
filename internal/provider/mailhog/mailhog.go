package mailhog

import (
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
	if msgTo := msg.Header.Get(mailutils.HeaderTo); msgTo != "" {
		to = append(to, msgTo)
	}

	return smtp.SendMail(Addr, nil, From, to, data)
}
