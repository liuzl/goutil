package rest

import (
	"encoding/json"
	"net/http"
)

type RestMessage struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type Message struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ErrBadRequest(w http.ResponseWriter, message interface{}) {
	MustEncodeWithStatus(w,
		map[string]*Message{"error": {http.StatusBadRequest, message}},
		http.StatusBadRequest)
}

func ErrInternalServer(w http.ResponseWriter, message interface{}) {
	MustEncodeWithStatus(w,
		map[string]*Message{"error": {http.StatusInternalServerError, message}},
		http.StatusInternalServerError)
}

func ErrMethodNotAllowed(w http.ResponseWriter, message interface{}) {
	MustEncodeWithStatus(w,
		map[string]*Message{"error": {http.StatusMethodNotAllowed, message}},
		http.StatusMethodNotAllowed)
}

func ErrorMessageWithStatus(w http.ResponseWriter, message interface{}, status int) {
	MustEncodeWithStatus(w,
		map[string]*Message{"error": {status, message}},
		status)
}

func MustEncode(w http.ResponseWriter, i interface{}) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		e.Encode(err.Error())
	}
}

func MustEncodeWithStatus(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	if err := e.Encode(i); err != nil {
		e.Encode(err.Error())
	}
}
