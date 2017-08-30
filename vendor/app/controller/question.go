package controller

import (
	"net/http"
	"app/db"
)

func init() {
	http.Handle("/q", ApiHandler(getQuestions))
	http.Handle("/q/total", ApiHandler(getQuestionsTotal))
}

func getQuestions(w http.ResponseWriter, r *http.Request) *ApiResult {
	list, err := db.GetQuestions()
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
