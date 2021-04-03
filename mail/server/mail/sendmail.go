package main

import (
	"github.com/go-mail/mail"
)

func (s *Server) sendMail(message *mail.Message) error {
	log := log("sendMail")

	recipient := message.GetHeader("To")[0]

	log.Tracef("Sending mail to %s", recipient)

	// send the mail
	return s.dialer.DialAndSend(message)
}
