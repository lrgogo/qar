package controller

import (
	"net/http"
	"app/db"
	"strconv"
	"app/util"
)

func init() {
	http.Handle("/u_i", ApiHandler(getUserInfo))
}

func getUserInfo(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	_, err := strconv.ParseInt(r.Form.Get("uid"), 10, 64)
	if err != nil {
		return util.Error(util.PARAMS_ERROR, "参数错误")
	}
	info, err := db.GetUserInfo(r.Form.Get("uid"))
	if err != nil {
		return err
	}
	return util.Success(info)
}
