package httputil

import (
	"net/http"
	"strconv"
)

func GetParamInt(r *http.Request, param string) int {
	get, ok := r.URL.Query()["seed"]
	if ok && len(get[0]) > 0 {
		out, err := strconv.Atoi(string(get[0]))
		if err != nil {
			panic(err)
		}
		return out
	}
	return 0
}

func GetParamStr(r *http.Request, param string) string {
	get, ok := r.URL.Query()["seed"]
	if ok && len(get[0]) > 0 {
		out := string(get[0])
		return out
	}
	return ""
}
