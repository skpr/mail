package ses

import (
	"context"
	"fmt"
	"log"
	"net/mail"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"

	"github.com/skpr/mail/internal/mailutils"
)

// AccessKeyPrefix is used when identifying if a credential is used for AWS IAM authentication.
// https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html
const AccessKeyPrefix = "AKIA"

// Send email via AWS SES.
func Send(region, username, password, from string, to []string, msg *mail.Message) error {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     username,
				SecretAccessKey: password,
				SessionToken:    "",
				Source:          "github.com/skpr/mail",
			},
		}),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatal(err)
	}

	if val, ok := msg.Header[mailutils.HeaderTo]; ok {
		to = append(to, val...)
	}
	msg.Header[mailutils.HeaderTo] = to

	err = mailutils.EnforceFrom(msg, from)
	if err != nil {
		return fmt.Errorf("failed to set from header: %w", err)
	}

	data, err := mailutils.MessageToBytes(msg)
	if err != nil {
		return fmt.Errorf("failed to convert message to bytes: %w", err)
	}

	input := &ses.SendRawEmailInput{
		RawMessage: &types.RawMessage{
			Data: data,
		},
		Source: aws.String(from),
	}

	output, err := ses.NewFromConfig(cfg).SendRawEmail(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to send message via ses %w", err)
	}

	log.Printf("successfully sent message via ses with id %s", *output.MessageId)

	return nil
}
