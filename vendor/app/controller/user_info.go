package controller

import (
	"net/http"
	"app/db"
	"strconv"
)

func init() {
	http.Handle("/u_i", ApiHandler(getUserInfo))
}

func getUserInfo(w http.ResponseWriter, r *http.Request) *ApiResult{
	if err := r.ParseForm(); err != nil {
		return &ApiResult{Error: err}
	}
	_, err := strconv.ParseInt(r.Form.Get("uid"), 10, 64)
	if err != nil {
		return &ApiResult{Code: CODE_PARAMS_ERROR, Msg:"参数错误"}
	}
	info, err := db.GetUserInfo(r.Form.Get("uid"))
	if err != nil {
		return &ApiResult{Error: err}
	}
	return &ApiResult{Code: CODE_SUCCESS, Data: info}
}
