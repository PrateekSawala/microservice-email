package main

import (
	"flag"
	"net/http"
	"os"

	"microsevice.email/mail/rpc/mail"

	"github.com/sirupsen/logrus"
)

var (
	service           = flag.String("service", os.Getenv("SERVICE_NAME"), "Service name")
	SmtpAccountEmail  = flag.String("SMTP account email", os.Getenv("SMTP_Account_Email"), "The email account of mail sender")
	serverFQDNandPort = "localhost:3015"
)

func main() {

	/* Logging */
	logrus.SetLevel(logrus.TraceLevel)

	log := log("main")

	// Init Mail connection
	goMailDialer := Client()

	/* Server */

	// Init server context
	server := Server{dialer: *goMailDialer}

	handler := mail.NewMailServer(&server, nil)

	// Check if port is provided in environment configuration
	if os.Getenv("PORT") != "" {
		serverFQDNandPort = os.Getenv("PORT")
	}

	log.Debugf("Microsevice %s server started: %v", *service, serverFQDNandPort)

	log.Warnf("Server exited: %s", http.ListenAndServe(serverFQDNandPort, handler))
}
