package mailutils

import (
	"bytes"
	"fmt"
	"net/mail"

	qp "gopkg.in/alexcesaro/quotedprintable.v3"
)

const (
	// HeaderFRom is used to identify the From header for an email.
	HeaderFrom = "From"

	// The term CRLF refers to Carriage Return (ASCII 13, \r) Line Feed (ASCII 10, \n).
	// They're used to note the termination of a line.
	crlf = "\r\n"
)

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
