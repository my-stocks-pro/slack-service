package main

import (
	"net/http"
	"time"
)

type Server interface {
	New(r TypeRouter) TypeServer
}

type TypeServer struct {
	HTTP *http.Server
}

func NewServer(router Router) TypeServer {
	return TypeServer{
		HTTP: &http.Server{
			Handler:      router.InitMux(),
			Addr:         ":8002",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func (s TypeServer) Serve() error {
	return s.HTTP.ListenAndServe()
}
