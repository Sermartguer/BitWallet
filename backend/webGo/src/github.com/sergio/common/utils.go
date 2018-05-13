package common

import (
	"os"
	"reflect"
	"unsafe"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//SendMail used to send an email
func SendMail(typeSend string, id string, username string, userMail string) {
	from := mail.NewEmail("BitWallet", "bitwallet@bitwallet.com")
	if typeSend == "check" {
		subject := "Account Verification"
		to := mail.NewEmail(username, userMail)
		plainTextContent := "asd"
		stringComplete := "http://localhost:5000/verify/" + id
		htmlContent := "<div>To verify, <a href='" + stringComplete + "'>click here</a></div>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		client.Send(message)

	} else if typeSend == "newPassword" {
		subject := "Account Verification"
		to := mail.NewEmail(username, userMail)
		plainTextContent := "asd"
		stringComplete := "http://localhost:5000/newpassword/" + id
		htmlContent := "<div>To recover password, <a href='" + stringComplete + "'>click here</a></div>"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		client.Send(message)
	}

}

//BytesToString converts a []byte var to string
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
