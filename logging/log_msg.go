package logging

import "github.com/sirupsen/logrus"

var localLogger = logrus.New()

func Info(msg ...interface{}) {
	localLogger.Info(msg)
}

func Debug(msg ...interface{}) {
	localLogger.Debug(msg)
}

func Error(err ...interface{}) {
	localLogger.Error(err)
}
