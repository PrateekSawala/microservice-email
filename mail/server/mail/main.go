package main

import (
	"fmt"
	"net/http"

	"microsevice.email/mail/rpc/mail"

	"github.com/sirupsen/logrus"
)

var (
	service      = "mail"           // service name
	serverPort   = "localhost:8080" // Server port
	smtpMail     = ""               // SMTP account email
	smptUser     = ""               // SMTP account user
	smtpPASSWORD = ""               // SMTP account password
)

func main() {

	/* Logging */
	logrus.SetLevel(logrus.TraceLevel)

	log := log("main")

	fmt.Println("main")

	// Init Mail connection
	goMailDialer := Client()

	/* Server */

	// Init server context
	server := Server{dialer: *goMailDialer}

	handler := mail.NewMailServer(&server, nil)

	log.Debugf("Microsevice.email %s server started: %v", service, serverPort)

	log.Warnf("Server exited: %s", http.ListenAndServe(serverPort, handler))
}