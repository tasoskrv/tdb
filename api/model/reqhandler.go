package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func DecodeBody(r *http.Request, v interface{}) error {
	if r.Body != nil {
		defer r.Body.Close()
		return json.NewDecoder(r.Body).Decode(v)
	}
	return errors.New("Undefined body")
}

func EncodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if data != nil {
		EncodeBody(w, r, data)
	}
}

func RespondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

func RespondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	RespondErr(w, r, status, http.StatusText(status))
}
