package logger

import (
	"os"
	"github.com/sirupsen/logrus"
)

// LogClass - Logger class
type LogClass struct {
	Logger  *logrus.Logger
}

// NewInitiateLogger - constructor
func NewInitiateLogger() *LogClass {
	lc := &LogClass{
		Logger: logrus.New(),
	}

	lc.Logger.Formatter = &logrus.JSONFormatter{}
	lc.Logger.Out = os.Stdout
	lc.Logger.SetLevel(logrus.TraceLevel)
	return lc
}

// Log - log
func (lc *LogClass) Log() *logrus.Logger {
	return lc.Logger
}
