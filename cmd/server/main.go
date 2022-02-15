package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"src/pkg/api"
	"src/pkg/common/helpers"
	"src/pkg/common/middleware"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Ujuzi API is Healthy"))
}

func run() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.Use(middleware.LoggingMiddleware)

	helpers.Mount(router, "/api/v1", api.Router())

	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(srv.ListenAndServe())
}

func main() {
	run()
}
