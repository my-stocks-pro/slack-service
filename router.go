package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router interface {
	New(Logger TypeLogger) TypeRouter
}

type TypeRouter struct {
	Logger TypeLogger
	Router *mux.Router
}

func NewRouter() TypeRouter {
	return TypeRouter{}
}

func (r TypeRouter) New(Logger TypeLogger) TypeRouter {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", r.HandlerHealth).Methods(http.MethodGet)
	router.HandleFunc("/slack", r.Slack).Methods(http.MethodPost)
	return TypeRouter{Router: router}
}

func (r TypeRouter) Slack(w http.ResponseWriter, req *http.Request) {

}
