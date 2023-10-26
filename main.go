package main

import (
	"errors"
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/alecthomas/kingpin/v2"

	skprconfig "github.com/skpr/go-config"
	"github.com/skpr/mail/internal/provider/mailhog"
	"github.com/skpr/mail/internal/provider/ses"
)

const (
	// ConfigKeyUsername used when authenticating with SMTP on the Skpr hosting platform.
	ConfigKeyUsername = "smtp.username"
	// ConfigKeyPassword used when authenticating with SMTP on the Skpr hosting platform.
	ConfigKeyPassword = "smtp.password"
	// ConfigKeyRegion used when authenticating with SMTP on the Skpr hosting platform.
	ConfigKeyRegion = "smtp.region"
	// ConfigKeyFromAddress used for overriding the FROM address.
	ConfigKeyFromAddress = "smtp.from.address"
)

var (
	to                = kingpin.Arg("to", "The list of recipients separated by comma.").Strings()
	from              = kingpin.Flag("from", "The from address (ignored)").Short('f').String()
	recipientsFromMsg = kingpin.Flag("to-from-message", "Read message for to (ignored)").Short('t').Bool()
	ignoreDots        = kingpin.Flag("ignore-dots", "Ignore dots alone on lines by themselves in incoming messages (ignored).").Short('i').Bool()
)

func main() {
	kingpin.Parse()

	if *from != "" {
		log.Println("Ignoring flag -f", *from)
	}
	if *recipientsFromMsg == true {
		log.Println("Ignoring flag -t")
	}
	if *ignoreDots == true {
		log.Println("Ignoring flag -i")
	}

	config, err := skprconfig.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	var (
		username = config.GetWithFallback(ConfigKeyUsername, os.Getenv("SKPRMAIL_USERNAME"))
		password = config.GetWithFallback(ConfigKeyPassword, os.Getenv("SKPRMAIL_PASSWORD"))
		region   = config.GetWithFallback(ConfigKeyRegion, os.Getenv("SKPRMAIL_REGION"))
		from     = config.GetWithFallback(ConfigKeyFromAddress, os.Getenv("SKPRMAIL_FROM"))
	)

	msg, err := mail.ReadMessage(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message from stdin: %s", err)
	}

	err = send(region, username, password, from, *to, msg)
	if err != nil {
		log.Fatalf("failed to send message: %s", err)
	}
}

// Send email based on parameters.
func send(region, username, password, from string, to []string, msg *mail.Message) error {
	// Use AWS if the credentials match what we would expect for IAM.
	if strings.HasPrefix(username, ses.AccessKeyPrefix) {
		return ses.Send(region, username, password, from, to, msg)
	}

	return mailhog.Send(to, msg)
}
