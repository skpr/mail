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
 To: my@email.com
 Subject: sendmail test two
 From: me@myserver.com
 
 And here goes the e-mail body, test test test..
 ```
 
 ```bash
export SKPRMAIL_AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
export SKPRMAIL_AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}
export SKPRMAIL_AWS_REGION=us-east-1
export SKPRMAIL_FROM_ADDRESS=admin@previousnext.com.au
 $ cat /tmp/test-email.txt | skprmail
 ```
 