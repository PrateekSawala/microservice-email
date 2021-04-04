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
	// sendTestMail()
	// sendCustomMail()
}

func sendTestMail() {

	// Declare mail content
	mailContent := &mail.SendTestMailInput{
		Name:    "",
		Email:   "",
		Message: "",
		Phone:   "",
	}

	_, err := client.SendTestMail(context.Background(), mailContent)
	if err != nil {
		log.Println("Error", err)
	}
}

func sendCustomMail() {
	// Declare mail content
	mailContent := &mail.SendCustomMailInput{
		From:    "",
		To:      "",
		Subject: "",
		Body:    "",
	}

	_, err := client.SendCustomMail(context.Background(), mailContent)
	if err != nil {
		log.Println("Error", err)
	}
}
