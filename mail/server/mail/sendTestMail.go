package main

import (
	"bytes"
	"context"
	"html/template"
	"io/ioutil"

	"microsevice.email/common/constant"
	"microsevice.email/mail/rpc/mail"
)

// SendTestMail ...
func (s *Server) SendTestMail(ctx context.Context, input *mail.SendTestMailInput) (*mail.SendTestMailResult, error) {
	log := log("SendTestMail")
	log.Tracef("Start")
	defer log.Tracef("End")

	// Check inputs
	if input == nil || input.Email == "" {
		return nil, constant.EmptyField
	}

	log.Tracef("Input: %+v", input)

	// Creating a new message
	m := s.NewMessage()

	//Setting headers
	m.SetHeader("From", *SmtpAccountEmail)
	m.SetHeader("To", input.Email)
	m.SetHeader("Subject", "Welcome")

	// init new input.Email content buffer
	content := new(bytes.Buffer)
	templateIntByte, err := ioutil.ReadFile("/templates/template.html")
	if err != nil {
		log.Debugf("Error while reading template file: %s", err)
		return nil, err
	}

	templateInt := template.Must(template.New("input.EmailTemplate").Parse(string(templateIntByte)))

	// Selecting mail greeting
	paramsBody := map[string]string{
		"Name":    input.Name,
		"Phone":   input.Phone,
		"Message": input.Message,
	}

	//Set body content of the email template
	body := template.Must(template.New("Template-Body").Parse(`<h3>Hey there,</h3><p>Name: {{.Name}}</p><p>Phone: {{.Phone}}</p><p>Message: {{.Message}}</p>`))

	bodyContent := new(bytes.Buffer)
	// Set parameters which will be filled in template
	body.Execute(bodyContent, &paramsBody)

	paramsTemplate := map[string]interface{}{
		"Body":    template.HTML(string(bodyContent.Bytes())),
		"Preview": input.Preview,
		"Title":   input.Title,
	}
	// Parse template and substitute params
	templateInt.Execute(content, &paramsTemplate)

	//Setting the body of the message
	m.SetBody("text/html", string(bodyContent.Bytes()))

	err = s.sendMail(m)
	if err != nil {
		log.Debugf("Error while sending internal mail to service@: %s", err)
		return nil, err
	}

	return &mail.SendTestMailResult{}, nil
}
