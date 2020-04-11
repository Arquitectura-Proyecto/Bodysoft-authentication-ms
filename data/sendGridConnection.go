package data

import (
	"fmt"

	"github.com/jpbmdev/Bodysoft-authentication-ms/credentials"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail ..
func SendEmail(email string, htmlContentp string, subjectp string) string {
	from := mail.NewEmail("BodySoft", "bodysoftms@gmail.com")
	to := mail.NewEmail("User", email)
	subject := subjectp
	plainTextContent := "servicio de autentificacion"
	htmlContent := htmlContentp
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(credentials.Sendgrid)
	response, err := client.Send(message)
	if err != nil {
		return err.Error()
	}
	fmt.Println(response.StatusCode)
	return "ok"
}
