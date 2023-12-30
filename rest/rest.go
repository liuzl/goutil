package rest

import (
	"encoding/json"
	"net/http"
)

type RestMessage struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrBadRequest(w http.ResponseWriter, message string) {
	MustEncodeWithStatus(w,
		map[string]*ErrorMessage{"error": {http.StatusBadRequest, message}},
		http.StatusBadRequest)
}

func ErrInternalServer(w http.ResponseWriter, message string) {
	MustEncodeWithStatus(w,
		map[string]*ErrorMessage{"error": {http.StatusInternalServerError, message}},
		http.StatusInternalServerError)
}

func ErrMethodNotAllowed(w http.ResponseWriter, message string) {
	MustEncodeWithStatus(w,
		map[string]*ErrorMessage{"error": {http.StatusMethodNotAllowed, message}},
		http.StatusMethodNotAllowed)
}

func ErrorMessageWithStatus(w http.ResponseWriter, message string, status int) {
	MustEncodeWithStatus(w,
		map[string]*ErrorMessage{"error": {status, message}},
		status)
}

func MustEncode(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		//panic(err)
		e.Encode(err.Error())
	}
}

func MustEncodeWithStatus(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		//panic(err)
		e.Encode(err.Error())
	}
}
