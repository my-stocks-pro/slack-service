package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router interface {
	InitMux() *mux.Router
}

type TypeRouter struct {
	Logger Logger
	Slack  Slack
}

func NewRouter(logger Logger, slack Slack) TypeRouter {
	return TypeRouter{
		Logger: logger,
		Slack:  slack,
	}
}

func (r TypeRouter) InitMux() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", r.HandlerHealth).Methods(http.MethodGet)
	router.HandleFunc("/slack", r.SlackHandler).Methods(http.MethodPost)
	return router
}

func (r TypeRouter) SlackHandler(w http.ResponseWriter, req *http.Request) {

}
