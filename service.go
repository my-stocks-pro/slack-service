package main

import "fmt"

type TypeService struct {
	Logger TypeLogger
	Slack  TypeSlack
	Router TypeRouter
}

func NewService(logger Logger, slack Slack, router Router) TypeService {

	log, err := logger.InitLogger()
	if err != nil {
		fmt.Println(err.Error())
	}

	return TypeService{
		Logger: log,
		Slack:  slack.New(log),
		Router: router.New(log),
	}
}
