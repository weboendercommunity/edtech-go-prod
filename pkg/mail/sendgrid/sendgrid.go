package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	registerDto "edtech.id/internal/register/dto"
)

type Mail interface {
	SendVerificationEmail(toEmail string, registerDto registerDto.CreateEmailVerification)
}

type MailImpl struct {
}

func (mi *MailImpl) sendMail(toEmail string, result string, subject string) {

	from := mail.NewEmail(os.Getenv("SENDGRID_MAIL_SENDER"), os.Getenv("SENDGRID_MAIL_SENDER"))
	to := mail.NewEmail(toEmail, toEmail)

	message := mail.NewSingleEmail(from, subject, to, "", result)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(message)

	if err != nil {
		fmt.Println(err)
	} else if response.StatusCode != 200 {
		fmt.Println(response)
	} else {
		fmt.Printf("Email sent to %s", toEmail)
	}

}

func (mi *MailImpl) SendVerificationEmail(toEmail string, registerDto registerDto.CreateEmailVerification) {
	cwd, _ := os.Getwd()

	templateFile := filepath.Join(cwd, "/templates/emails/verification_email.html")

	result, err := ParseTemplate(templateFile, registerDto)

	if err != nil {
		fmt.Println(err)
	}

	mi.sendMail(toEmail, result, registerDto.SUBJECT)
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	errParse := t.Execute(buf, data)

	if errParse != nil {
		return "", errParse
	}

	return buf.String(), nil
}

func NewMail() Mail {
	return &MailImpl{}
}
