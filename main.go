package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/skpr/go-skprconfig"
	extensionsv1beta1 "github.com/skpr/operator/pkg/apis/extensions/v1beta1"

	"github.com/skpr/mail/internal/mailhog"
	"github.com/skpr/mail/internal/ses"
)

func main() {
	var (
		username = skprconfig.Get(extensionsv1beta1.ConfigMapKeyUsername, os.Getenv("SKPRMAIL_USERNAME"))
		password = skprconfig.Get(extensionsv1beta1.SecretKeyPassword, os.Getenv("SKPRMAIL_PASSWORD"))
		region   = skprconfig.Get(extensionsv1beta1.ConfigMapKeyRegion, os.Getenv("SKPRMAIL_REGION"))
		from     = skprconfig.Get(extensionsv1beta1.ConfigMapKeyFromAddress, os.Getenv("SKPRMAIL_FROM"))
	)

	// Load input from Stdin and build up raw mail object.
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message from stdin: %s", err)
	}

	err = send(region, username, password, from, data)
	if err != nil {
		log.Fatalf("failed to send: %s", err)
	}
}

// Send email based on parameters.
func send(region, username, password, from string, data []byte) error {
	// Use AWS if the credentials match what we would expect for IAM.
	if strings.HasPrefix(username, ses.AccessKeyPrefix) {
		return ses.Send(region, username, password, from, data)
	}

	return mailhog.Send(data)
}
