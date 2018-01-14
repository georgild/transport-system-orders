package main

import (
	"fmt"
	"log"
	"net/smtp"
)

// SendMail :
func SendMail(mailTo string) {

	if len(mailTo) <= 0 {
		return
	}
	auth := smtp.PlainAuth("", "georgild77@gmail.com", "", "smtp.gmail.com")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{mailTo}
	msg := []byte("From: transport@system.org\r\n" +
		fmt.Sprintf("To: %s", mailTo) + "\r\n" +
		"Subject: Your order has been approved!\r\n" +
		"\r\n" +
		"Your order has been approved!\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, "transport@system.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
