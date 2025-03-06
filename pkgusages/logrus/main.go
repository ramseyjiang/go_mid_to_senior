package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	log.WithFields(logrus.Fields{
		"user":   "john.doe",
		"action": "login",
	}).Info("User logged in")
}
