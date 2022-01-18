package helpers

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendEmail(targetEmail, subject, name, url string) error {
	errEnv := godotenv.Load()
	if errEnv != nil {
		return  errors.New("Failed to load env file")
	}
	mailPort, _ := strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	var (
		ConfigSmtpHost =    os.Getenv("CONFIG_SMTP_HOST")
		ConfigSmtpPort =     mailPort
		ConfigSenderName =   os.Getenv("CONFIG_SENDER_NAME")
		ConfigAuthEmail  =    os.Getenv("CONFIG_AUTH_EMAIL")
		ConfigAuthPassword = os.Getenv("CONFIG_AUTH_PASSWORD")
		BaseUrlTarget = os.Getenv("BASE_URL_TARGET")
	)
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", ConfigSenderName)
	mailer.SetHeader("To", targetEmail)
	if subject == "register" {
		mailer.SetHeader("Subject", "Registration Account")
		mailer.SetBody("text/html", fmt.Sprintf("<h1 style='margin-top:0;margin-bottom:16px;font-size:26px;line-height:32px;font-weight:bold;letter-spacing:-0.02em;'>Welcome to Sewakantor!</h1></br>Hello <b> %s </b>\nPlease click this <a href=\"https://%s/%s\">link</a> to activate your account.</br> Thank you.", name,BaseUrlTarget,url))
	}


	dialer := gomail.NewDialer(
		ConfigSmtpHost,
		ConfigSmtpPort,
		ConfigAuthEmail,
		ConfigAuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}