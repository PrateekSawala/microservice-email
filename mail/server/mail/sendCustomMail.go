package main

import (
	"bytes"
	"context"
	"html/template"
	"io/ioutil"

	"microsevice.email/common/constant"
	"microsevice.email/mail/rpc/mail"
)

// SendCustomMail ...
func (s *Server) SendCustomMail(ctx context.Context, input *mail.SendCustomMailInput) (*mail.SendCustomMailResult, error) {
	log := log("SendCustomMail")
	log.Tracef("Start")
	defer log.Tracef("End")

	// Check inputs
	if input == nil || input.From == "" || input.To == "" || input.Subject == "" || input.Body == "" {
		log.Debugf("Empty Input")
		return nil, constant.EmptyField
	}

	// Creating a new message
	m := s.NewMessage()

	//Setting headers
	m.SetHeader("From", input.From)
	m.SetHeader("To", input.To)
	m.SetHeader("Subject", input.Subject)

	// init new input.Email content buffer
	content := new(bytes.Buffer)
	templateIntByte, err := ioutil.ReadFile("/templates/template.html")
	if err != nil {
		log.Debugf("Error while reading template file: %s", err)
		return nil, err
	}

	templateInt := template.Must(template.New("input.EmailTemplate").Parse(string(templateIntByte)))

	paramsBody := map[string]string{}

	//Set body content of the email template
	body := template.Must(template.New("Template-Body").Parse(input.Body))

	bodyContent := new(bytes.Buffer)

	// Set parameters which will be filled in template
	body.Execute(bodyContent, &paramsBody)

	paramsTemplate := map[string]interface{}{
		"Body":    template.HTML(string(bodyContent.Bytes())),
		"Preview": input.Preview,
		"Title":   input.Title,
	}

	if input.Title == "" {
		paramsTemplate["Title"] = input.Subject
	}

	// Parse template and substitute params
	templateInt.Execute(content, &paramsTemplate)

	//Setting the body of the message
	m.SetBody("text/html", string(content.Bytes()))

	//Calling the gatekeeper send mail method
	err = s.sendMail(m)
	if err != nil {
		log.Debugf("Error while sending  mail %s : to email : %s", err, input.To)
		return nil, err
	}

	return &mail.SendCustomMailResult{}, nil
}
