package controller

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"time"
)

type CommonResult struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

type ApiError struct {
	Error error
	Code  int
	Msg   string
}

type ApiHandler func(response http.ResponseWriter, r *http.Request) *ApiError

func (fn ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("2000-01-02 15:04:05"), r.Method, r.URL)

	if err := fn(w, r); err != nil {
		if err.Error != nil {
			log.Println(err)
			j, _ := json.Marshal(&CommonResult{
				Code: 500,
				Msg:  "server error",
			})

			w.Write(j)
		} else {
			j, _ := json.Marshal(&CommonResult{
				Code: err.Code,
				Msg:  err.Msg,
			})

			w.Write(j)
		}
	}

}

func Load()  {

}