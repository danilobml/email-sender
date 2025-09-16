# E-Mail Sender - A simple GO project to send emails

## Usage:

### Create a .env file with the variables:

- FROM_EMAIL=
- FROM_EMAIL_PASSWORD=
- FROM_EMAIL_SMTP="smtp.gmail.com"
- SMTP_ADDR="smtp.gmail.com:587"

(for Gmail)

### Run app:
```bash
go run ./cmd/api
```

### Send POST request:
to: http://localhost:8080/mail

body:
```json
{
  "to": []string,
  "subject": string,
  "body": string
}
```
