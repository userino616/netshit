package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	e       *logrus.Entry
	logFile *os.File
)

type Logger struct {
	*logrus.Entry
}

func GetLogger() Logger {
	return Logger{e}
}

func Init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{}) //nolint:exhaustivestruct

	err := os.MkdirAll("logs", 0o744)
	if err != nil {
		panic(err)
	}

	logFile, err = os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.MultiWriter(os.Stdout, logFile))

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}

func Close() error {
	return logFile.Close() //nolint:wrapcheck
}
