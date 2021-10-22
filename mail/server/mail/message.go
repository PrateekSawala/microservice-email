package main

import (
	"os"
	"strconv"

	"github.com/go-mail/mail"
)

// New Message
func (s *Server) NewMessage() *mail.Message {
	// Init new message
	return mail.NewMessage()
}

// Client
func Client() *mail.Dialer {
	// FInd SMTP port
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	// Init Mail connection
	d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	d.StartTLSPolicy = mail.MandatoryStartTLS
	d.Timeout = 20
	return d
}

// Server implements the mail service
type Server struct {
	dialer mail.Dialer
}
