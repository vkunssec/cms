package main

import (
	logs "cms/pkg/core"
	"cms/pkg/handler"
	"cms/pkg/middleware"

	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	uri := fmt.Sprintf("%s/?%s", os.Getenv("MONGODB_URI"), os.Getenv("MONGO_PARAM"))
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logs.Error.Fatalln(err)
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logs.Error.Fatalln(err)
		panic(err)
	}
	logs.Info.Println("connected mongodb -> ", uri)

	routers := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Logging),
	)

	handler.HandlerIndex(routers, *n, client)
	handler.HandlerArticle(routers, *n, client)

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
		panic(err)
	}

	logs.Info.Println("listen port", portDefault)
}
