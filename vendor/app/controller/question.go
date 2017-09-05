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

func getQuestions(w http.ResponseWriter, r *http.Request) error {
	uid, err := util.VerifyJWT(r)
	if err != nil {
		return err
	}
	list, err := db.GetQuestions(uid)
	if err != nil {
		return err
	}
	return util.Success(list)
}

func getQuestionsTotal(w http.ResponseWriter, r *http.Request) error {
	total, err := db.GetQuestionsTotal()
	if err != nil {
		return err
	}
	return util.Success(total)
}
