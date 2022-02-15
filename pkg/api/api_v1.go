package api

import (
	"github.com/gorilla/mux"
	"src/pkg/app/user"
	"src/pkg/common/helpers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	helpers.Mount(router, "/users", user.Router())

	return router
}
