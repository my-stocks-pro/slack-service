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
	HTTP *http.Server
}

func NewServer(logger TypeLogger) TypeServer {
	return TypeServer{
		Logger: logger,
		HTTP: &http.Server{}}
}

func (s TypeServer) InitServer(router Router) TypeServer {
	return TypeServer{
		HTTP: &http.Server{
			Handler:      router.InitMux().Router,
			Addr:         ":8002",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func (s TypeServer) Serve () error {
	if err:= s.HTTP.ListenAndServe(); err != nil {
		return err
	}
	return nil
}