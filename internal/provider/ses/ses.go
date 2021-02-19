package ses

import (
	"log"
	"net/mail"

	"github.com/skpr/mail/internal/mailutils"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// AccessKeyPrefix is used when identifying if a credential is used for AWS IAM authentication.
// https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html
const AccessKeyPrefix = "AKIA"

// Send email via AWS SES.
func Send(region, username, password, from string, to []string, msg *mail.Message) error {
	config := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(username, password, ""),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return err
	}

	if val, ok := msg.Header[mailutils.HeaderTo]; ok {
		to = append(to, val...)
	}
	msg.Header[mailutils.HeaderTo] = to

	err = mailutils.EnforceFrom(msg, from)
	if err != nil {
		return err
	}

	data, err := mailutils.MessageToBytes(msg)
	if err != nil {
		return err
	}

	input := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{
			Data: data,
		},
		Source: aws.String(from),
	}

	output, err := ses.New(sess).SendRawEmail(input)
	if err != nil {
		log.Fatalf("failed to send message: %s", err)
	}

	log.Printf("message id %s", *output.MessageId)

	return nil
}
