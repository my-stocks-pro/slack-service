package main

import (
	"fmt"
	"os"
	"go.uber.org/zap"
)

const (
	logPath   = "app_log"
	logFile   = "logs"
	logPrefix = "slack-service"
)

type Logger interface {
	InitLogger() (TypeLogger, error)
}

type TypeLogger struct {
	Client *zap.Logger
}

func NewLogger() TypeLogger {
	return TypeLogger{}
}

func (l TypeLogger) InitLogger() (TypeLogger, error) {
	filename := fmt.Sprintf("%s/%s_%s.log", logPath, logPrefix, logFile)
	_, err := os.Create(filename)
	if err != nil {
		return TypeLogger{}, err
	}

	conf := zap.NewDevelopmentConfig()
	conf.OutputPaths = []string{
		filename,
	}

	zapLog, err := conf.Build()
	if err != nil {
		return TypeLogger{}, err
	}

	defer zapLog.Sync()

	return TypeLogger{Client: zapLog}, nil
}

func (l TypeLogger) Error(err error) {
	l.Client.Error(err.Error())
}
