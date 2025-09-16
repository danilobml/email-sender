# E-Mail Sender - A simple GO project to send emails

## Usage:

### Clone repository
```bash
git clone https://github.com/danilobml/email-sender
```

### Install dependencies
```bash
cd email-sender
go mod tidy
```

### Create a .env file with the variables:

```txt
FROM_EMAIL=
FROM_EMAIL_PASSWORD=
FROM_EMAIL_SMTP="smtp.gmail.com"
SMTP_ADDR="smtp.gmail.com:587"
```
(Example for Gmail)

Then, add your email address to the first variable, get the password (e.g. in Gmail, go to account settings -> app passwords) and add to the second.

### Run app:
```bash
go run ./cmd/api
```

### Send POST request:
url: http://localhost:8080/mail

with json body:
```json
{
  "to": []string,
  "subject": string,
  "body": string
}
```
