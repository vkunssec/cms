package handler

import (
	"cms/pkg/core/entity"
	"cms/pkg/middleware"

	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
)

func handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var start time.Time

		latency := time.Since(start).Round(time.Second)

		response := entity.ResponseIndex{
			Version:     "1.0.0",
			Uptime:      latency.String(),
			Environment: os.Getenv("STAGE"),
			Message:     "OK",
			Date:        time.Now().UTC().Format("2006-01-02T15:04:05.999Z"),
		}

		headers := middleware.GetHeaders()
		for key, values := range headers {
			for _, value := range values {
				w.Header().Set(key, value)
			}
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})
}

func HandlerIndex(r *mux.Router, n negroni.Negroni, _ *mongo.Client) {
	r.Handle("/", n.With(
		negroni.Wrap(handler()),
	)).Methods("GET").Name("Index")
}
