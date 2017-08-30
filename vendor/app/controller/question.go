package controller

import (
	"net/http"
	"app/db"
)

func init() {
	http.Handle("/q", ApiHandler(getQuestions))
}

func getQuestions(w http.ResponseWriter, r *http.Request) *ApiResult {
	list, err := db.GetQuestions()
	if err != nil {
		return &ApiResult{Error: err}
	}
	return &ApiResult{Code: CODE_SUCCESS, Data: list}
}

