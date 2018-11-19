package main

import "fmt"

type TypeService struct {
	Logger TypeLogger
	Slack  TypeSlack
	Router TypeRouter
	Server TypeServer
}

func NewService(logger Logger) TypeService {

	log, err := logger.InitLogger()
	if err != nil {
		fmt.Println(err.Error())
	}
	slack := NewSlack().NewSlackClient(log)
	roter := NewRouter(log).InitMux()
	server := NewServer(log).InitServer(roter)

	return TypeService{
		Logger: log,
		Slack:  slack,
		Router: roter,
		Server: server,
	}
}
