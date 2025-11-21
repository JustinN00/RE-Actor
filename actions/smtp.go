package actions

import (
	"net/smtp"
	"log/slog"
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
	byte_message := []byte(message)
	err := smtp.SendMail(
		mail.SMTPHost + ":" + mail.SMTPPort,
		auth,
		mail.SendFrom,
		[]string{mail.SendTo},
		byte_message,
	)
	if err != nil {
		slog.Error("Error or something")
	}
	return nil, err
}

func (mail SMTPMail) Act(message string) error {
	_, err := mail.SendMail(message)
	return err
}