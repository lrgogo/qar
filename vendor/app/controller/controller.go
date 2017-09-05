package controller

import (
	"net/http"
	"log"
	"encoding/json"
	"app/util"
)

type ApiHandler func(w http.ResponseWriter, r *http.Request) error

func (fn ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.RemoteAddr, r.Body, r.Form, r.RequestURI, r.Header)

	err := fn(w, r)
	if _, ok := err.(*util.Response); !ok {
		log.Println(err)
		err = &util.Response{
			Code: util.SERVER_ERROR,
			Msg:  "服务器错误",
		}
	}

	j, _ := json.Marshal(err)
	log.Println(r.Method, r.URL, r.RemoteAddr, string(j))
	w.Write(j)

}

func Load() {

}
