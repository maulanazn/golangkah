package logging_test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSetLevel(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("Trace")
	logger.Debug("Debugging")
	logger.Info("Infoooo")
}

func TestLogOutputFile(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	
	file, _ := os.OpenFile("application.log", os.O_CREATE|os.O_RDONLY|os.O_APPEND, 0666)
	logger.SetOutput(file)

	logger.Trace("IDK")
}

func TestLogFieldFile(t *testing.T) {
	logger := logrus.New()

	logger.SetLevel(logrus.TraceLevel)
	
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.WithFields(logrus.Fields{
		"action": "login",
		"username": "maulanazn",
	}).Info("IDK")
}

func TestLogEntryFile(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	entry := logrus.NewEntry(logger)

	entry.WithFields(logrus.Fields{
		"action": "product",
		"username": "maulanazn",
	}).Info()
}
