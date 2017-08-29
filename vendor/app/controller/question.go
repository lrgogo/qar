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
		return &ApiResult{err, 0, "", nil}
	}
	return &ApiResult{nil, 1, "", list}
}

