package controller

import (
	"net/http"
	"app/db"
	"app/util"
)

func init() {
	http.Handle("/q", ApiHandler(getQuestions))
	http.Handle("/q/total", ApiHandler(getQuestionsTotal))
}

func getQuestions(w http.ResponseWriter, r *http.Request) *ApiResult {
	uid, err := util.VerifyJWT(r)
	if err != nil {
		return &ApiResult{Error:err}
	}
	list, err := db.GetQuestions(uid)
	if err != nil {
		return &ApiResult{Error: err}
	}
	return &ApiResult{Code: CODE_SUCCESS, Data: list}
}

func getQuestionsTotal(w http.ResponseWriter, r *http.Request) *ApiResult {
	total, err := db.GetQuestionsTotal()
	if err != nil {
		 return &ApiResult{Error:err}
	}
	return &ApiResult{Code:CODE_SUCCESS, Data:total}
}
