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

func NewServer() TypeServer {
	return TypeServer{}
}

func (s TypeServer) New(r TypeRouter) TypeServer {
	return TypeServer{
		HTTP: &http.Server{
			Handler:      r.Router,
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