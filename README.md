# skprmail

This project is a drop-in replacement for Sendmail which sends email via Amazon SES. 

It reads a raw SMTP message from stdin and forwards it unaltered.

## Installation

Download the binary for the latest release and place it at `$(which sendmail)`.

## Credentials

### skpr config
The app knows how skpr mounts configuration into the container (via the mount at `/etc/skpr`). It picks up credentials from the following keys:

* IAM Key ID `smtp.username`
* IAM Key Secret `smtp.password`
* AWS Region `smtp.region`
* From Address `smtp.from.address`

### Environment variables

* IAM Key ID `SKPRMAIL_AWS_ACCESS_KEY_ID`
* IAM Key Secret `SKPRMAIL_AWS_SECRET_ACCESS_KEY`
* AWS Region `SKPRMAIL_AWS_REGION`
* From Address `SKPRMAIL_FROM_ADDRESS`
 
## Testing Locally
   
Create a file `/tmp/test-mail.txt` with the following contents. Adjust from/to address to those which are verified in SES.
 ```
FROM:me@myserver.com
To: my@email.com
Subject: sendmail test two

And here goes the e-mail body, test test test..
 ```
 
 ```bash
export SKPRMAIL_AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
export SKPRMAIL_AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}
export SKPRMAIL_AWS_REGION=us-east-1
export SKPRMAIL_FROM_ADDRESS=admin@previousnext.com.au
 $ cat /tmp/test-email.txt | skprmail
 ```

### Testing with local smtp server

- Create a fresh build of the binary using `make build`
- Create a file `/tmp/test-mail.txt` with the following contents. Adjust from/to address to those which are verified in SES.
    ```
    FROM:me@myserver.com
    To: my@email.com
    Subject: sendmail test two

    And here goes the e-mail body, test test test..
    ```
- `export SKPRMAIL_ADDR="localhost:1025"` (to point skprmail to our debug server)
- Start the smtp server in a new terminal window from `utils/smtp-debug-server/smtp-debug-server` (please note that the smtp server has a delay of 45s during startup, enabling us to test timeout on skprmail)
- Run `$ cat /tmp/test-email.txt | bin/skprmail_linux_amd64 --timeout 10s` to test a 10 second timeout or run `$ cat /tmp/test-email.txt | bin/skprmail_linux_amd64` for a default 30s timeout. (please run the test command within 10 seconds of starting the debug server)
- Timeout error `Contacting the local smtp server timed out, cancelling...` will confirm timeout was engaged. 
- Once the SMTP server starts after 45s with log message
    ```
    2024/07/25 07:29:21 Starting SMTP server at 127.0.0.1:1025
    220 localhost ESMTP Service Ready
    ```
    You can test delivery of the email.
    ```
    2024/07/25 08:05:06 Starting SMTP server at 127.0.0.1:1025
    220 localhost ESMTP Service Ready
    EHLO localhost
    250-Hello localhost
    250-PIPELINING
    250-8BITMIME
    250-ENHANCEDSTATUSCODES
    250-CHUNKING
    250 SIZE
    MAIL FROM:<mail@skpr.localhost> BODY=8BITMIME
    250 2.0.0 Roger, accepting mail from <mail@skpr.localhost>
    RCPT TO:<my@email.com>
    250 2.0.0 I'll make sure <my@email.com> gets this
    DATA
    354 Go ahead. End your data with <CR><LF>.<CR><LF>
    From: me@myserver.com
    To: my@email.com
    Subject: sendmail test two

    And here goes the e-mail body, test test test..
    .
    250 2.0.0 OK: queued
    QUIT
    221 2.0.0 Bye
    ```
 
## Releasing

Testing a release:

```
goreleaser release --rm-dist --snapshot --skip-publish
```

Releasing:

```
golreleaser
```



