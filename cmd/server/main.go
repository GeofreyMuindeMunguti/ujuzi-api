package main

import (
	"log"
	"net/http"
	"src/pkg/api"
	"strings"

	"github.com/gorilla/mux"
)

func index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("API is Healthy"))
}

func run() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")

	mount(router, "/api/v1", api.Router())

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

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
