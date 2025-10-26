package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)

	entry.WithField("username", "ucup")

	entry.Info("Hello Entry")
	entry.Warn("Hello Entry")
	entry.Error("Hello Entry")

}
