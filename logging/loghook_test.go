package logging_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Its error, sending email.......", entry.Level, entry.Message)
	return nil
}

func TestLogHookFile(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	file, _ := os.OpenFile("weirdwire-log-cupcake.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Tes info")	
	logger.Error("Tes error")	
	logger.Warn("Tes warn")
}
