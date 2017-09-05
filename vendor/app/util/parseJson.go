package util

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func ParseJson(r *http.Request) (map[string]string, error) {
	//解析JSON
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var kv map[string]string
	err = json.Unmarshal(body, &kv)
	if err != nil {
		return nil, err
	}
	return kv, nil
}
