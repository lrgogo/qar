package controller

import (
	"net/http"
	"app/db"
	"strconv"
	"app/util"
)

func init() {
	http.Handle("/a", ApiHandler(getAnswers))
}

func getAnswers(w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		return util.Error(util.PARAMS_ERROR, "参数错误")
	}
	count, err := strconv.Atoi(r.Form.Get("count"))
	if err != nil {
		return util.Error(util.PARAMS_ERROR, "参数错误")
	}

	list, err := db.GetAnswers(page, count)
	if err != nil {
		return err
	}
	return util.Success(list)
}
