package initalize

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// initializes the logger for the application
func InitLogger() {
	log = logrus.New()
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	log.SetLevel(logrus.InfoLevel)
}

// return the initialized logger
func GetLogger() *logrus.Logger {
	return log
}