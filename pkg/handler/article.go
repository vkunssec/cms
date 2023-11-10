package handler

import (
	"cms/pkg/core/entity"
	"cms/pkg/middleware"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
)

func getArticle(_ mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	// vars := mux.Vars(r)
	// articleId := vars["id"]
}

func saveArticle(client mongo.Client) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article entity.Article

		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// save article

		headers := middleware.GetHeaders()
		for key, values := range headers {
			for _, value := range values {
				w.Header().Set(key, value)
			}
		}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]interface{}{
			"saved":   true,
			"article": &article,
		})
	})
}

func HandlerArticle(r *mux.Router, n negroni.Negroni, client *mongo.Client) {
	r.Handle("/article/{id}", n.With(
		negroni.Wrap(
			getArticle(*client),
		),
	)).Methods("GET").Name("GetArticle")

	r.Handle("/article", n.With(
		negroni.Wrap(
			saveArticle(*client),
		),
	)).Methods("POST").Name("SaveArticle")
}
