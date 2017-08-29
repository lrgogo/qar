package controller

import (
	"net/http"
	"app/db"
	"encoding/json"
)

func init() {
	http.Handle("/q", ApiHandler(getQuestions))
}

func getQuestions(w http.ResponseWriter, r *http.Request) *ApiError {
	list, err := db.GetQuestions()
	j, err := json.Marshal(&CommonResult{
		Code: 1,
		Msg:  "",
		Data: list,
	})
	if err != nil {
		return &ApiError{err, 0, ""}
	}

	w.Write(j)
	return nil
}

