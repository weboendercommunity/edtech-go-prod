package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	registerDto "edtech.id/internal/register/dto"

	"gopkg.in/gomail.v2"
)

type SmtpMail interface {
	SmtpSendVerificationEmail(toEmail string, dto registerDto.CreateEmailVerification)
	SmtpSendMail(toEmail string, result string, subject string)
}

type SmtpMailImpl struct{}

func (mi *SmtpMailImpl) SmtpSendMail(toEmail string, result string, subject string) {
	msg := gomail.NewMessage()

	// Set sender
	msg.SetHeader("From", os.Getenv("SMTP_MAIL_SENDER"))

	// Set recipient
	msg.SetHeader("To", toEmail)

	// Set subject
	msg.SetHeader("Subject", subject)

	// Set body
	msg.SetBody("text/html", result)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		panic(err)
	}

	// Create a new Sendinblue sender with API key
	sender := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_MAIL_SENDER"), os.Getenv("SMTP_PASSWORD"))
	sender.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email
	// Send the email
	if err := sender.DialAndSend(msg); err != nil {
		panic(err)
	}
}

// SendVerificationEmail implements Mail
func (mi *SmtpMailImpl) SmtpSendVerificationEmail(toEmail string, dto registerDto.CreateEmailVerification) {
	cwd, _ := os.Getwd()
	templateFile := filepath.Join(cwd, "/templates/emails/verification_email.html")

	result, err := SmtpParseTemplate(templateFile, dto)

	if err != nil {
		fmt.Print(err)
	}

	mi.SmtpSendMail(toEmail, result, dto.SUBJECT)
}

func SmtpParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func NewSmtpMail() SmtpMail {
	return &SmtpMailImpl{}
}
