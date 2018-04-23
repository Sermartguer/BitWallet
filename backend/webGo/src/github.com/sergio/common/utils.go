package common

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func GetTokenParsed(t string) (jwt.MapClaims, bool) {
	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("do or do not there is no try"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
func SendMail(typeSend string, id string, username string, userMail string) {
	from := mail.NewEmail("BitWallet", "bitwallet@bitwallet.com")
	if typeSend == "check" {
		subject := "Account Verification"
		to := mail.NewEmail(username, userMail)
		plainTextContent := "asd"
		stringComplete := "http://localhost:5000/verify/" + id
		htmlContent := "<div>To verify, <a href='" + stringComplete + "'>click here</a></div>"
		fmt.Println(htmlContent)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	} else if typeSend == "newPassword" {
		subject := "Account Verification"
		to := mail.NewEmail(username, userMail)
		plainTextContent := "asd"
		stringComplete := "http://localhost:5000/newpassword/" + id
		htmlContent := "<div>To recover password, <a href='" + stringComplete + "'>click here</a></div>"
		fmt.Println(htmlContent)
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
		}
	}

}
