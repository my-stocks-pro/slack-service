package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
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
	dataType := req.Header.Get("type")

	switch dataType {
	case "earnings", "approved", "rejected":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			r.Logger.Error(err)
			http.Error(w, "Read body error", http.StatusBadRequest)
		}

		if err := r.Slack.Send(body, dataType); err != nil {
			r.Logger.Error(err)
			http.Error(w, "Slack service error", http.StatusServiceUnavailable)
		}
	default:
		r.Logger.Info("Type not found")
		http.Error(w, "Type not found", http.StatusNotFound)
	}
}
