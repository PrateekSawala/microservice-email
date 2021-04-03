package main

import (
	"github.com/go-mail/mail"
)

// New Message
func (s *Server) NewMessage() *mail.Message {
	// Init new message
	return mail.NewMessage()
}

// Client
func Client() *mail.Dialer {
	// Init Mail connection
	d := mail.NewDialer(smtpHostName, smtpPortNumber, smptUserName, smtpAccountPassword)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	return d
}

// Server implements the mail service
type Server struct {
	dialer mail.Dialer
}
