package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello logging")
	logger.Warn("Hello logging")
	logger.Error("Hello logging")
}
