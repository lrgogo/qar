package controller

import (
	"net/http"
)

func init() {
	http.Handle("/user", ApiHandler(getUsers))
}

func getUsers(w http.ResponseWriter, r *http.Request) *ApiResult {
	return &ApiResult{}
}
