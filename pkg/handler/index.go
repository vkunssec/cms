package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type ResponseIndex struct {
	Uptime      string `json:"uptime"`
	Message     string `json:"message"`
	Version     string `json:"version"`
	Environment string `json:"env"`
	Date        string `json:"date"`
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {
	var start time.Time

	latency := time.Since(start).Round(time.Second)

	response := ResponseIndex{
		Version:     "1.0.0",
		Uptime:      latency.String(),
		Environment: os.Getenv("STAGE"),
		Message:     "OK",
		Date:        time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
