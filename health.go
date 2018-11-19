package main

import (
	"net/http"
	"time"
	"encoding/json"
	"errors"
)

type HealthCheckType struct {
	Service  string
	CurrDate string
	Version  string
}

func (r TypeRouter) HandlerHealth(w http.ResponseWriter, req *http.Request) {
	data := HealthCheckType{
		Service:  "slack-service",
		CurrDate: time.Now().Format("2006-01-02 15:04"),
		Version:  "1.0",
	}

	r.Logger.Error(errors.New("HealthCheckType"))

	payload, err := json.Marshal(data)
	if err != nil {
		r.Logger.Error(errors.New("HandlerHealth Marshal"))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
