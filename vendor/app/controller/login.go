package controller

import (
	"net/http"
	"app/util"
	"app/db"
)

func init() {
	http.Handle("/login", ApiHandler(login))
}

func login(w http.ResponseWriter, r *http.Request) error {
	kv, err := util.ParseJson(r)
	if err != nil {
		return err
	}
	mobile := kv["mobile"]
	pwd := kv["pwd"]
	if mobile == "" || pwd == "" {
		return util.Error(util.PARAMS_ERROR, "参数错误")
	}
	info, err := db.GetLoginInfo(mobile, pwd)
	if err != nil {
		return err
	}
	token, err := util.CreateAccessJWT(info.Id)
	if err != nil {
		return err
	}
	info.Token = token
	return util.Success(info)
}

