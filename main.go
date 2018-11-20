package main

import "fmt"

const ServiceName = "slack-service"

func main() {

	logger, err := NewLogger()
	if err != nil {
		fmt.Println(err.Error())
	}

	slack := NewSlack()

	router := NewRouter(logger, slack)

	server := NewServer(router)

	logger.Info(fmt.Sprintf("Start %s on port %s", ServiceName, server.HTTP.Addr))

	logger.Error(server.Serve())
}
