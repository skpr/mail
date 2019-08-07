package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kelseyhightower/envconfig"

	"github.com/skpr/mail/internal/skprconfig"
)

const (
	ConfigKeyAccessId     string = "smtp.username"
	ConfigKeyAccessSecret string = "smtp.password"
	ConfigKeyRegion       string = "smtp.region"
	ConfigKeyFrom         string = "smtp.from.address"
)

type Params struct {
	ConfigBasePath     string `envconfig:"config_base_path" default:"/etc/skpr"`
	AwsAccessKeyId     string `envconfig:"aws_access_key_id"`
	AwsSecretAccessKey string `envconfig:"aws_secret_access_key"`
	AwsRegion          string `envconfig:"aws_region"`
	FromAddress        string `envconfig:"from_address"`
}

func main() {
	// Load app behavior from environment.
	var params Params
	err := envconfig.Process("skprmail", &params)
	if err != nil {
		log.Fatal("Could not process configuration")
	}

	// Check for skpr config values for any parameters in Params which are empty.
	c := skprconfig.NewConfig(params.ConfigBasePath)
	if params.AwsAccessKeyId == "" {
		params.AwsAccessKeyId, err = c.Get(ConfigKeyAccessId)
		if err != nil {
			log.Fatal("AWS credentials not configured")
		}
	}
	if params.AwsSecretAccessKey == "" {
		params.AwsSecretAccessKey, err = c.Get(ConfigKeyAccessSecret)
		if err != nil {
			log.Fatal("AWS credentials not configured")
		}
	}
	if params.AwsRegion == "" {
		params.AwsRegion, err = c.Get(ConfigKeyRegion)
		if err != nil {
			log.Fatal("AWS region not configured")
		}
	}
	if params.FromAddress == "" {
		params.FromAddress, err = c.Get(ConfigKeyFrom)
		if err != nil {
			log.Println("FROM address not configured. This may impact deliverability of the message.")
		}
	}

	sess, err := awsSession(params)
	if err != nil {
		// @todo fallback to smtp forwarding (i.e. mailhog) if aws session not available.
		log.Fatal("AWS region not configured")
	}
	client := ses.New(sess)

	// Load input from Stdin and build up raw mail object.
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message from stdin: %s", err)
	}
	mailinput := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{Data: input},
		Source:     &params.FromAddress,
	}

	// Send email.
	output, err := client.SendRawEmail(mailinput)
	if err != nil {
		log.Fatalf("failed to send message: %s", err)
	}

	// Everything looks OK!
	log.Printf("message id %s", *output.MessageId)
	os.Exit(0)
}

// Helper function which creates AWS session.
func awsSession(params Params) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:      aws.String(params.AwsRegion),
		Credentials: credentials.NewStaticCredentials(params.AwsAccessKeyId, params.AwsSecretAccessKey, ""),
	})
}
