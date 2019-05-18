package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

//GetLogger ... Will return the logging object
func GetLogger(envType string) logrus.Logger {
	var logger logrus.Logger

	if envType == "production" {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.TextFormatter{})

		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		logger.SetOutput(os.Stdout)

		// Only log the warning severity or above.
		logger.SetLevel(logrus.InfoLevel)

		return logger
	} else if envType == "development" {
		// Log as JSON instead of the default ASCII formatter.
		logger.SetFormatter(&logrus.TextFormatter{})

		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		logger.SetOutput(os.Stdout)

		// Only log the warning severity or above.
		logger.SetLevel(logrus.DebugLevel)

		return logger
	}
	return logger
}
