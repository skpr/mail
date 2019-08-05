package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kelseyhightower/envconfig"

	"github.com/skpr/mail/internal/skprconfig"
)

type Params struct {
	ConfigBasePath     string `default:"/etc/skpr"`
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsRegion          string
	FromAddress        string
}

func main() {
	// Load app behavior from environment.
	var params Params
	err := envconfig.Process("skprmail", &params)
	if err != nil {
		fmt.Println("Could not process configuration")
		os.Exit(1)
	}

	// Check for skpr config values for any parameters in Params which are empty.
	c := skprconfig.NewConfig(params.ConfigBasePath, skprconfig.DefaultTrimSuffix)
	if params.AwsAccessKeyId == "" {
		params.AwsAccessKeyId, err = c.GetWithError("smtp.username")
		if err != nil {
			fmt.Println("AWS credentials not configured")
			os.Exit(1)
		}
	}
	if params.AwsSecretAccessKey == "" {
		params.AwsSecretAccessKey, err = c.GetWithError("smtp.password")
		if err != nil {
			fmt.Println("AWS credentials not configured")
			os.Exit(1)
		}
	}
	if params.AwsRegion == "" {
		params.AwsRegion, err = c.GetWithError("smtp.region")
		if err != nil {
			fmt.Println("AWS region not configured")
			os.Exit(1)
		}
	}
	if params.FromAddress == "" {
		params.FromAddress, err = c.GetWithError("smtp.region")
		if err != nil {
			fmt.Println("FROM address not configured. This may impact deliverability of the message.")
		}
	}

	sess, err := awsSession(params)
	if err != nil {
		fmt.Println("Could not initialise AWS session.")
		os.Exit(1)
	}
	client := ses.New(sess)

}

// Helper function which creates AWS session.
func awsSession(params Params) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Region:      aws.String(params.AwsRegion),
		Credentials: credentials.NewStaticCredentials(params.AwsAccessKeyId, params.AwsSecretAccessKey, ""),
	})
}
