package controller

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"time"
)

type ApiResult struct {
	Error error `json:"error,omitempty"`
	Code  int `json:"code"`
	Msg   string `json:"msg"`
	Data  interface{} `json:"data,omitempty"`
}

type ApiHandler func(response http.ResponseWriter, r *http.Request) *ApiResult

func (fn ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("2000-01-02 15:04:05"), r.Method, r.URL)

	result := fn(w, r)
	if result.Error != nil {
		log.Println(result.Error)
	}
	j, _ := json.Marshal(result)

	w.Write(j)

}

func Load() {

}
