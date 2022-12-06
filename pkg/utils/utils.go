package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if b, err := ioutil.ReadAll(r.Body); err == nil {
		//json 문자열을 byte[]로 변환
		if err := json.Unmarshal([]byte(b), x); err != nil {
			return
		}
	}
}
