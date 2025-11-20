package actions

import (
	"net/smtp"
)

type SMTPMail struct {
	SMTPHost string
	SMTPPort string
	SendFrom string
	SendTo string
	Password string
}

func (mail SMTPMail) SendMail(message string) (error, error) {
	auth := smtp.PlainAuth("", mail.SendFrom, mail.Password, mail.SMTPHost)
	byte_message = []byte(message)
	error := smtp.SendMail(
		mail.SMTPHost + ":" + mail.SMTPPort,
		auth,
		mail.SendFrom,
		mail.SendTo,
		byte_message,
	)
	if err != nil {
		slog.Error("Error or something")
	}
	return "", error
}

func (mail SMTPMail) Act(message string) error {
	_, error := dwh.SendMail(message)
	return err
}