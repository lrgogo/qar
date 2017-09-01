package controller

import (
	"net/http"
	"app/db"
	"strconv"
)

func init() {
	http.Handle("/a", ApiHandler(getAnswers))
}

func getAnswers(w http.ResponseWriter, r *http.Request) *ApiResult{
	if err := r.ParseForm(); err != nil {
		return &ApiResult{Error: err}
	}
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		return &ApiResult{Code: CODE_PARAMS_ERROR, Msg:"参数错误"}
	}
	count, err := strconv.Atoi(r.Form.Get("count"))
	if err != nil {
		return &ApiResult{Code: CODE_PARAMS_ERROR, Msg:"参数错误"}
	}

	list, err := db.GetAnswers(page, count)
	if err != nil {
		return &ApiResult{Error: err}
	}
	return &ApiResult{Code: CODE_SUCCESS, Data: list}
}