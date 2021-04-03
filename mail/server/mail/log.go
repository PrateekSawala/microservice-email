package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

// log is a helper class that enrichens the structured logging
func log(input interface{}) *logrus.Entry {
	log := logrus.WithFields(logrus.Fields{
		"service": service,
		"method":  "unkown",
	})

	// Check input
	if input == nil {
		return log
	}

	// Check input type
	switch input.(type) {
	case string:
		return log.WithField("method", input)
	default:
		// add raw input
		inputType := strings.Split(fmt.Sprintf("%v", reflect.TypeOf(input)), ".")
		method := inputType[len(inputType)-1]
		method = strings.TrimSuffix(method, "Input")
		rawInput := fmt.Sprintf("%+v", input)
		if rawInput != "<nil>" {
			log = log.WithField("input", rawInput)
		}
		return log.WithField("method", method)
	}
}
