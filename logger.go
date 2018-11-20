package main

import (
	"fmt"
	"os"
	"go.uber.org/zap"
)

const (
	logPath   = "app_log"
	logPrefix = "slack-service"
)

type Logger interface {
	Error(err error)
	Info(msg string)
}

type TypeLogger struct {
	Client *zap.Logger
}

func NewLogger() (TypeLogger, error) {
	filename := fmt.Sprintf("%s/%s.log", logPath, logPrefix)
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

func (l TypeLogger) Info(msg string) {
	l.Client.Info(msg)
}
