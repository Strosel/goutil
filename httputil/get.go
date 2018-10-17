package httputil

import (
	"errors"
	"net/http"
	"strconv"
)

func GetParamInt(r *http.Request, param string) (int, error) {
	get, ok := r.URL.Query()[param]
	if ok && len(get[0]) > 0 {
		out, err := strconv.Atoi(string(get[0]))
		return out, err
	}
	return 0, errors.New("No parameter named " + param)
}

func GetParamFloat(r *http.Request, param string) (float64, error) {
	get, ok := r.URL.Query()[param]
	if ok && len(get[0]) > 0 {
		out, err := strconv.ParseFloat(string(get[0]), 64)
		return out, err
	}
	return 0, errors.New("No parameter named " + param)
}

func GetParamString(r *http.Request, param string) string {
	get, ok := r.URL.Query()[param]
	if ok && len(get[0]) > 0 {
		out := string(get[0])
		return out
	}
	return ""
}
