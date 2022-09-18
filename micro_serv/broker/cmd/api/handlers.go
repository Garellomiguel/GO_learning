package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool        `json:error`
	Message string      `json:message`
	Data    interface{} `json:data,omitempty`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hi the broker",
	}

	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
