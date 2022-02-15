package user

import (
	"net/http"
)

func listUsers(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Users"))
}
