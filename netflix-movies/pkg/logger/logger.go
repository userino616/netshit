package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	e       *logrus.Entry
)

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func Init(logLvl logrus.Level) {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})

	l.SetOutput(os.Stdout)
	l.SetLevel(logLvl)
	e = logrus.NewEntry(l)
}