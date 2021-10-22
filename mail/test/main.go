package main

import (
	"context"
	"fmt"
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
		Name:    "test user",
		Email:   "testuser@mail.com",
		Message: "test message",
		Phone:   "1234567890",
	}
	// SendTestMail
	_, err := client.SendTestMail(context.Background(), mailContent)
	if err != nil {
		log.Println("Error", err)
	}
}

func sendCustomMail() {
	// Declare mail content
	mailContent := &mail.SendCustomMailInput{
		From:    "testuser@mail.com",
		To:      "testuser1@mail.com",
		Subject: "test message",
	}
	// Prepare content body
	mailContent.Body = fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	    <head>
			<title>Test message</title>
		</head>
	    <body>
			<p>Hi there</p>
		</body>
	</html>
	`)
	// SendCustomMail
	_, err := client.SendCustomMail(context.Background(), mailContent)
	if err != nil {
		log.Println("Error", err)
	}
}
