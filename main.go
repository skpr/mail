package main

import (
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/skpr/go-skprconfig"
	extensionsv1beta1 "github.com/skpr/operator/pkg/apis/extensions/v1beta1"

	"github.com/skpr/mail/internal/provider/mailhog"
	"github.com/skpr/mail/internal/provider/ses"
)

func main() {
	var (
		username = skprconfig.Get(extensionsv1beta1.ConfigMapKeyUsername, os.Getenv("SKPRMAIL_USERNAME"))
		password = skprconfig.Get(extensionsv1beta1.SecretKeyPassword, os.Getenv("SKPRMAIL_PASSWORD"))
		region   = skprconfig.Get(extensionsv1beta1.ConfigMapKeyRegion, os.Getenv("SKPRMAIL_REGION"))
		from     = skprconfig.Get(extensionsv1beta1.ConfigMapKeyFromAddress, os.Getenv("SKPRMAIL_FROM"))
	)

	msg, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message: %s", err)
	}

	err = send(region, username, password, from, msg)
	if err != nil {
		log.Fatalf("failed to send: %s", err)
	}
}

// Send email based on parameters.
func send(region, username, password, from string, msg *mail.Message) error {
	// Use AWS if the credentials match what we would expect for IAM.
	if strings.HasPrefix(username, ses.AccessKeyPrefix) {
		return ses.Send(region, username, password, from, msg)
	}

	return mailhog.Send(msg)
}
