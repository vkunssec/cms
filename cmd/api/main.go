package main

import (
	logs "cms/pkg/core"
	"cms/pkg/handler"
	"cms/pkg/middleware"

	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	var portDefault = 3000
	portEnv, ok := os.LookupEnv("PORT")
	if !ok {
		os.Setenv("PORT", "3000")
		port, _ := strconv.Atoi(os.Getenv("PORT"))
		portDefault = port
	} else {
		portDefault, _ = strconv.Atoi(portEnv)
	}

	routers := mux.NewRouter()

	routers.Use(middleware.Logging)
	routers.Use(middleware.Cors)

	routers.HandleFunc("/", handler.IndexHandler).
		Methods("GET").
		Name("index")

	http.Handle("/", routers)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", portDefault),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routers,
	}
	if err := srv.ListenAndServe(); err != nil {
		logs.Error.Fatalln(err)
	}

	logs.Info.Println("listen port", portDefault)
}
