package controller

import (
	"net/http"
	"log"
	"encoding/json"
	"time"
)

const (
	CODE_SUCCESS      = 1
	CODE_SERVER_ERROR = 2
)

type ApiResult struct {
	Error error `json:"-"`
	Code  int `json:"code"`
	Msg   string `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
}

type ApiHandler func(response http.ResponseWriter, r *http.Request) *ApiResult

func (fn ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL, r.RemoteAddr, r.Body, r.Form, r.RequestURI, r.Header)

	result := fn(w, r)
	if result.Error != nil {
		log.Println(result.Error)
		result.Code = CODE_SERVER_ERROR
		result.Msg = "服务器错误"
	}
	j, _ := json.Marshal(result)
	log.Println(time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL, r.RemoteAddr, string(j))
	w.Write(j)

}

func Load() {

}
