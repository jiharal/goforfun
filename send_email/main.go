package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

const (
	Email    = "Email"
	Password = "Haha"
)

func main() {
	SentByEmail("jihar1997@gmail.com", "Verification code", "2132 Is the verification code for displae application")
}

func SentByEmail(To, Subject, Body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", Email)
	m.SetHeader("To", To)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", Body)
	d := gomail.NewDialer("smtp.gmail.com", 587, Email, Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
