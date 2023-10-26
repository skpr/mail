package mailutils

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnforceFrom(t *testing.T) {
	scenarios := []struct {
		Message  *mail.Message
		From     string
		Expected *mail.Message
	}{
		// No From address was specified.
		{
			Message: &mail.Message{
				Header: mail.Header{
					HeaderTo: []string{
						"to@example.com",
					},
				},
				Body: strings.NewReader("This is a test"),
			},
			From: "override@example.com",
			Expected: &mail.Message{
				Header: mail.Header{
					HeaderTo: []string{
						"to@example.com",
					},
					HeaderFrom: []string{
						"override@example.com",
					},
					HeaderSender: []string{
						"override@example.com",
					},
					HeaderReplyTo: []string{
						"override@example.com",
					},
				},
				Body: strings.NewReader("This is a test"),
			},
		},
		// No Reply-To header was set.
		{
			Message: &mail.Message{
				Header: mail.Header{
					HeaderTo: []string{
						"to@example.com",
					},
					HeaderFrom: []string{
						"from@example.com",
					},
				},
				Body: strings.NewReader("This is a test"),
			},
			From: "override@example.com",
			Expected: &mail.Message{
				Header: mail.Header{
					HeaderTo: []string{
						"to@example.com",
					},
					HeaderFrom: []string{
						"override@example.com",
					},
					HeaderSender: []string{
						"override@example.com",
					},
					HeaderReplyTo: []string{
						"from@example.com",
					},
				},
				Body: strings.NewReader("This is a test"),
			},
		},
		// Don't override existing Reply-To header.
		{
			Message: &mail.Message{
				Header: mail.Header{
					HeaderTo: []string{
						"to@example.com",
					},
					HeaderFrom: []string{
						"from@example.com",
					},
					HeaderReplyTo: []string{
						"alreadyoverridden@example.com",
					},
				},
				Body: strings.NewReader("This is a test"),
			},
			From: "override@example.com",
			Expected: &mail.Message{
				Header: mail.Header{
					HeaderTo: []string{
						"to@example.com",
					},
					HeaderFrom: []string{
						"override@example.com",
					},
					HeaderSender: []string{
						"override@example.com",
					},
					HeaderReplyTo: []string{
						"alreadyoverridden@example.com",
					},
				},
				Body: strings.NewReader("This is a test"),
			},
		},
	}

	for _, scenario := range scenarios {
		err := EnforceFrom(scenario.Message, scenario.From)
		assert.Nil(t, err)
		assert.Equal(t, scenario.Expected.Header, scenario.Message.Header)
	}
}

func TestMessageToBytes(t *testing.T) {
	msg := &mail.Message{
		Header: mail.Header{
			HeaderTo: []string{
				"test@example.com",
			},
			HeaderFrom: []string{
				"test@example.com",
			},
		},
		Body: strings.NewReader("This is a test"),
	}

	b, err := MessageToBytes(msg)
	assert.Nil(t, err)

	assert.Contains(t, string(b), "To: test@example.com")
	assert.Contains(t, string(b), "From: test@example.com")
	assert.Contains(t, string(b), "This is a test")
}
