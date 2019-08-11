package logging

import "github.com/sirupsen/logrus"

var localLogger = logrus.New()

func Info(msg string) {
	localLogger.Info(msg)
}

func Debug(msg string) {
	localLogger.Debug(msg)
}

func Error(err error) {
	localLogger.Error(err)
}
