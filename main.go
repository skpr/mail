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
	"github.com/skpr/go-skprconfig"
	extensionsv1beta1 "github.com/skpr/operator/pkg/apis/extensions/v1beta1"
)

// Params which are loaded from environment variables.
type Params struct {
	ConfigBasePath     string `envconfig:"config_base_path" default:"/etc/skpr"`
	AwsAccessKeyID     string `envconfig:"aws_access_key_id"`
	AwsSecretAccessKey string `envconfig:"aws_secret_access_key"`
	AwsRegion          string `envconfig:"aws_region"`
	FromAddress        string `envconfig:"from_address"`
}

func main() {
	var params Params

	// Load app behavior from environment.
	err := envconfig.Process("skprmail", &params)
	if err != nil {
		log.Fatal("Could not process configuration")
	}

	// Check for skpr config values for any parameters in Params which are empty.
	c := skprconfig.NewConfig(params.ConfigBasePath)

	var (
		username = c.GetWithFallback(extensionsv1beta1.ConfigMapKeyUsername, params.AwsAccessKeyID)
		password = c.GetWithFallback(extensionsv1beta1.SecretKeyPassword, params.AwsSecretAccessKey)
		region   = c.GetWithFallback(extensionsv1beta1.ConfigMapKeyRegion, params.AwsRegion)
		address  = c.GetWithFallback(extensionsv1beta1.ConfigMapKeyFromAddress, params.FromAddress)
	)

	if username == "" {
		log.Fatal("AWS credentials not configured")
	}

	if password == "" {
		log.Fatal("AWS credentials not configured")
	}

	if region == "" {
		log.Fatal("AWS region not configured")
	}

	if address == "" {
		log.Println("FROM address not configured. This may impact deliverability of the message.")
	}

	config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(username, password, ""),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		// @todo fallback to smtp forwarding (i.e. mailhog) if aws session not available.
		log.Fatal("AWS region not configured")
	}

	// Load input from Stdin and build up raw mail object.
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read message from stdin: %s", err)
	}

	input := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{Data: data},
		Source:     &params.FromAddress,
	}

	// Send email.
	output, err := ses.New(sess).SendRawEmail(input)
	if err != nil {
		log.Fatalf("failed to send message: %s", err)
	}

	// Everything looks OK!
	log.Printf("message id %s", *output.MessageId)
	os.Exit(0)
}
