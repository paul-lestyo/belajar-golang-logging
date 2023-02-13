package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "pauls").Info("Hello World")

	logger.WithField("username", "pauls").
		WithField("name", "Paulus Lestyo").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "pauls",
		"name":     "Paulus Lestyo A",
	}).Info("Hello World")
}
