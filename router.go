package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router interface {
	InitMux() TypeRouter
}

type TypeRouter struct {
	Logger TypeLogger
	Router *mux.Router
}

func NewRouter(logger TypeLogger) TypeRouter {
	return TypeRouter{
		Logger: logger,
		Router: mux.NewRouter().StrictSlash(true)}
}

func (r TypeRouter) InitMux() TypeRouter {
	r.Router.HandleFunc("/health", r.HandlerHealth).Methods(http.MethodGet)
	r.Router.HandleFunc("/slack", r.Slack).Methods(http.MethodPost)
	return TypeRouter{Router: r.Router}
}

func (r TypeRouter) Slack(w http.ResponseWriter, req *http.Request) {

}
