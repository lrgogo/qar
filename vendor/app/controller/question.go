package controller

import (
	"net/http"
	"app/db"
	"encoding/json"
)

func init() {
	http.HandleFunc("/q", getQuestions)
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	list := db.GetQuestions()
	j, err := json.Marshal(&CommonResult{
		Code: 1,
		Msg:  "",
		Data: list,
	})
	checkErr(err)

	w.Write(j)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
