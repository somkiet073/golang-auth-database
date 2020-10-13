package jsonify

import (
	"encoding/json"
	"net/http"
)

// func response json
// @param1 http.ResponseWriter
// @param2 code Response example Http Status 200
// @param3 interface example { "response":"OK" }
func Json(w http.ResponseWriter) func(int, interface{}) error {
	return func(code int, v interface{}) error {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(code)
		return json.NewEncoder(w).Encode(v)
	}
}

// func request json
// @param1 http.Request
// @param2 interface example body
func Bind(r http.Request) func(interface{}) error {
	return func(v interface{}) error {
		return json.NewDecoder(r.Body).Decode(v)
	}
}
