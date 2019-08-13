package mailutils

import (
	"net/mail"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	expected := "To: test@example.com\r\nFrom: test@example.com\r\n\r\nThis is a test"

	assert.Equal(t, expected, string(b))
}
