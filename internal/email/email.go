package email

import (
	"fmt"
	"os"
	"gopkg.in/gomail.v2"
)

func SendMagicLink(to string, token string) error{
	link := fmt.Sprintf("%s/auth/verify?token=%s", os.Getenv("APP_URL"), token)

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("Email_USER"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your Login Link (Valid for 5 minutes)")
	m.SetBody("text/plain", "Click the link to log in:\n\n"+link)

	d := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		587,
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASS"),
	)

	return d.DialAndSend(m)
}