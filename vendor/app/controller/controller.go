package controller

import (
	"net/http"
	"log"
	"encoding/json"
)

const (
	CODE_SUCCESS      = 1
	CODE_SERVER_ERROR = 2
	CODE_PARAMS_ERROR = 3
)

type ApiResult struct {
	Error error `json:"-"`
	Code  int `json:"code"`
	Msg   string `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
}

type ApiHandler func(w http.ResponseWriter, r *http.Request) *ApiResult

func (fn ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.RemoteAddr, r.Body, r.Form, r.RequestURI, r.Header)

	result := fn(w, r)
	if result.Error != nil {
		log.Println(result.Error)
		result.Code = CODE_SERVER_ERROR
		result.Msg = result.Error.Error()
	}
	j, _ := json.Marshal(result)
	log.Println(r.Method, r.URL, r.RemoteAddr, string(j))
	w.Write(j)

}

func Load() {

}
