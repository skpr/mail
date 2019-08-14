package mailutils

import (
	"bytes"
	"fmt"
	"net/mail"

	qp "gopkg.in/alexcesaro/quotedprintable.v3"
)

const (
	// HeaderFrom is used to identify the From header for an email.
	HeaderFrom = "From"
	// HeaderReplyTo is used to identify the Reply-To header for an email.
	HeaderReplyTo = "Reply-To"
	// HeaderSender is used to identify the Sender header for an email.
	HeaderSender = "Sender"
	// HeaderTo is used to identify the To header for an email.
	HeaderTo = "To"

	// The term CRLF refers to Carriage Return (ASCII 13, \r) Line Feed (ASCII 10, \n).
	// They're used to note the termination of a line.
	crlf = "\r\n"
)

// EnforceFrom ensures that all outgoing mail comes from the correct FROM while preserving intent in the Reply-To.
func EnforceFrom(msg *mail.Message, from string) error {
	var reply []string

	if val, ok := msg.Header[HeaderFrom]; ok {
		reply = val
	} else {
		reply = []string{from}
	}

	// Set Reply-To as the origin From address if Reply-To is not set.
	if _, ok := msg.Header[HeaderReplyTo]; !ok {
		msg.Header[HeaderReplyTo] = reply
	}

	// These need to be set.
	msg.Header[HeaderFrom] = []string{from}
	msg.Header[HeaderSender] = []string{from}

	return nil
}

// MessageToBytes converts a Message to a set of bytes ready for delivery.
// Inspired by https://github.com/mohamedattahri/mail/blob/master/message.go#L263
func MessageToBytes(msg *mail.Message) ([]byte, error) {
	raw := &bytes.Buffer{}

	for key, items := range msg.Header {
		for _, item := range items {
			if item != "" {
				_, err := fmt.Fprintf(raw, "%s: %s%s", key, qp.QEncoding.Encode("utf-8", item), crlf)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	raw.WriteString(crlf)

	_, err := raw.ReadFrom(msg.Body)
	if err != nil {
		return nil, err
	}

	return raw.Bytes(), nil
}
