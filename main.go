package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	kingpin "github.com/alecthomas/kingpin/v2"

	skprconfig "github.com/skpr/go-config"
	defaultprovider "github.com/skpr/mail/internal/provider/default"
	"github.com/skpr/mail/internal/provider/ses"
)

const (
	// EnvUsername used to configure authentication using an environment variable.
	EnvUsername = "SKPRMAIL_USERNAME"
	// EnvPassword used to configure authentication using an environment variable.
	EnvPassword = "SKPRMAIL_PASSWORD"
	// EnvRegion used to configure the region where SES resides using an environment variable.
	EnvRegion = "SKPRMAIL_SES_REGION"
	// EnvFrom used to override and set the FROM request for mail using an environment variable.
	EnvFrom = "SKPRMAIL_FROM"
	// EnvHostname is the hostname to dispatch the email to.
	EnvHostname = "SKPRMAIL_HOSTNAME"
	// EnvPort is the port to be used to dispatch the email.
	EnvPort = "SKPRMAIL_PORT"
	// ConfigUsername used to configure authentication using a Skpr config.
	ConfigUsername = "smtp.username"
	// ConfigPassword used to configure authentication using a Skpr config.
	ConfigPassword = "smtp.password"
	// ConfigRegion used to configure the region where SES resides using a Skpr config.
	ConfigRegion = "smtp.region"
	// ConfigFrom used to override and set the FROM request for mail using a Skpr config.
	ConfigFrom = "smtp.from.address"
	// ConfigHostname used to override and set the address where mail is sent using a Skpr config.
	ConfigHostname = "smtp.hostname"
	// ConfigPort used to override and set the address port where mail is sent using a Skpr config.
	ConfigPort = "smtp.port"
)

var (
	cliTo                = kingpin.Arg("to", "The list of recipients separated by comma.").Strings()
	cliFrom              = kingpin.Flag("from", "The from address (ignored)").Short('f').String()
	cliRecipientsFromMsg = kingpin.Flag("to-from-message", "Read message for to (ignored)").Short('t').Bool()
	cliIgnoreDots        = kingpin.Flag("ignore-dots", "Ignore dots alone on lines by themselves in incoming messages (ignored).").Short('i').Bool()
	cliTimeout           = kingpin.Flag("timeout", "How long to wait before timing out and exiting.").Default("30s").Duration()
)

func main() {

	kingpin.Parse()

	if *cliFrom != "" {
		log.Println("Ignoring flag -f", *cliFrom)
	}

	if *cliRecipientsFromMsg {
		log.Println("Ignoring flag -t")
	}

	if *cliIgnoreDots {
		log.Println("Ignoring flag -i")
	}

	config, err := skprconfig.Load()
	if err != nil && !errors.Is(err, skprconfig.ErrNotFound) {
		panic(err)
	}

	var (
		username = config.GetWithFallback(ConfigUsername, os.Getenv(EnvUsername))
		password = config.GetWithFallback(ConfigPassword, os.Getenv(EnvPassword))
		region   = config.GetWithFallback(ConfigRegion, os.Getenv(EnvRegion))
		from     = config.GetWithFallback(ConfigFrom, os.Getenv(EnvFrom))
		hostname = config.GetWithFallback(ConfigHostname, os.Getenv(EnvHostname))
		port     = config.GetWithFallback(ConfigPort, os.Getenv(EnvPort))
	)

	msg, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message from stdin: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), *cliTimeout)
	defer cancel()

	err = send(ctx, region, username, password, hostname, port, from, *cliTo, msg)
	if err != nil {
		log.Fatalf("failed to send message: %s", err)
	}
}

// Send email based on parameters.
func send(ctx context.Context, region, username, password, hostname, port, from string, to []string, msg *mail.Message) error {
	// Use AWS if the credentials match what we would expect for IAM.
	if strings.HasPrefix(username, ses.AccessKeyPrefix) {
		return ses.Send(ctx, region, username, password, from, to, msg)
	}

	destination := fmt.Sprintf("%s:%s", hostname, port)

	return defaultprovider.Send(ctx, destination, to, msg)
}
