package controller

import (
	"net/http"
	"app/util"
	"app/db"
)

func init() {
	http.Handle("/login", ApiHandler(login))
}

func login(w http.ResponseWriter, r *http.Request) *ApiResult{
	kv, err := util.ParseJson(r)
	mobile := kv["mobile"]
	pwd := kv["pwd"]
	if mobile == "" || pwd == ""{
		return &ApiResult{Code:CODE_PARAMS_ERROR, Msg:"参数错误"}
	}
	info, err := db.GetLoginInfo(mobile, pwd)
	if err != nil {
		return &ApiResult{Error:err}
	}
	token, err := util.CreateAccessJWT(info.Id)
	if err != nil {
		return &ApiResult{Error:err}
	}
	info.Token = token
	return &ApiResult{Code:CODE_SUCCESS, Data:info}
}

