package belajar_golang_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello logger")
	fmt.Println("Hello logger")
}
