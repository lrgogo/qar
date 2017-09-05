package controller

import (
	"net/http"
	"app/util"
)

func init() {
	http.Handle("/user", ApiHandler(getUsers))
}

func getUsers(w http.ResponseWriter, r *http.Request) error {
	return util.Success("")
}
