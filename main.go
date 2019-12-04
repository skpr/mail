package main

import (
	"log"
	"net/mail"
	"os"
	"strings"
	"errors"

	skprconfig "github.com/skpr/go-config"
	extensionsv1beta1 "github.com/skpr/operator/pkg/apis/extensions/v1beta1"

	"github.com/skpr/mail/internal/provider/mailhog"
	"github.com/skpr/mail/internal/provider/ses"
)

func main() {
	config, err := skprconfig.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist){
		panic(err)
	}

	var (
		username = config.GetWithFallback(extensionsv1beta1.ConfigMapKeyUsername, os.Getenv("SKPRMAIL_USERNAME"))
		password = config.GetWithFallback(extensionsv1beta1.SecretKeyPassword, os.Getenv("SKPRMAIL_PASSWORD"))
		region   = config.GetWithFallback(extensionsv1beta1.ConfigMapKeyRegion, os.Getenv("SKPRMAIL_REGION"))
		from     = config.GetWithFallback(extensionsv1beta1.ConfigMapKeyFromAddress, os.Getenv("SKPRMAIL_FROM"))
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
