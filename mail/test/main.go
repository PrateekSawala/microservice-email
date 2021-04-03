package main

import (
	"context"
	"log"
	"net/http"

	"microsevice.email/mail/rpc/mail"
)

var (
	client mail.Mail
)

func main() {
	client = mail.NewMailProtobufClient("http://localhost:8080", &http.Client{})
	sendTestMail()
}

func sendTestMail() {
	_, err := client.SendTestMail(context.Background(), &mail.SendTestMailInput{})
	if err != nil {
		log.Println("Error", err)
	}
}
