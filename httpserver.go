package main

import (
	"net/http"
	"time"
)

type HTTPServer interface {
	InitServer(router Router) TypeServer
}

type TypeServer struct {
	Logger TypeLogger
	Server *http.Server
}

func NewServer(logger TypeLogger) TypeServer {
	return TypeServer{
		Logger: logger,
		Server: &http.Server{}}
}

func (s TypeServer) InitServer(router Router) TypeServer {
	return TypeServer{
		Server: &http.Server{
			Handler:      router.InitMux().Router,
			Addr:         ":8002",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}
