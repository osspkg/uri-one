package logger

import (
	"os"
	"uri-one/app/config"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	file *os.File
}

func MustNew(cfg *config.Config) *Logger {

	var (
		file *os.File
		err  error
	)

	if cfg.IsDebug() {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		file, err = os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			panic(err)
		} else {
			logrus.SetOutput(file)
		}
	}

	logrus.SetLevel(logrus.DebugLevel)

	return &Logger{file: file}
}

func (l *Logger) Start() error {
	logrus.Info("logger is started")

	return nil
}

func (l *Logger) Stop() error {
	logrus.Info("logger is stoped")

	if l.file != nil {
		if err := l.file.Close(); err != nil {
			panic(err)
		}
	}

	return nil
}
