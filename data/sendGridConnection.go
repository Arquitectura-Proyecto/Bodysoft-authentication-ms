package data

import (
	"fmt"

	"github.com/jpbmdev/Bodysoft-authentication-ms/localdata"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendEmail ..
func SendEmail(email string, pass string) string {
	from := mail.NewEmail("BodySoft", "bodysoftms@gmail.com")
	to := mail.NewEmail("User", email)
	subject := "Recuperacion de Contraseña Bodysoft"
	plainTextContent := "servicio de autentificacion"
	htmlContent := "<h1> Su contraseña es: " + pass + "</h1>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(localdata.Sendgrid)
	response, err := client.Send(message)
	if err != nil {
		return err.Error()
	}
	fmt.Println(response.StatusCode)
	return "ok"
}
