package controller

import (
	"net/http"
	"app/db"
)

func init() {
	http.Handle("/u_i", ApiHandler(getUserInfo))
}

func getUserInfo(w http.ResponseWriter, r *http.Request) *ApiResult{
	if err := r.ParseForm(); err != nil {
		return &ApiResult{Error: err}
	}
	uid := r.Form.Get("uid")
	info, err := db.GetUserInfo(uid)
	if err != nil {
		return &ApiResult{Error: err}
	}
	return &ApiResult{Code: CODE_SUCCESS, Data: info}
}
