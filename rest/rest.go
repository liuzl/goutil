package rest

import (
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

func GetClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		if i := strings.Index(ip, ","); i != -1 {
			return strings.TrimSpace(ip[:i])
		}
		return ip
	}
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}
	return r.RemoteAddr
}

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

func ErrNotFound(w http.ResponseWriter, message interface{}) {
	MustEncodeWithStatus(w,
		map[string]*Message{"error": {http.StatusNotFound, message}},
		http.StatusNotFound)
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
	e.SetEscapeHTML(false)
	if err := e.Encode(i); err != nil {
		e.Encode(err.Error())
	}
}

func MustEncodeWithStatus(w http.ResponseWriter, i interface{}, status int) {
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	if err := e.Encode(i); err != nil {
		e.Encode(err.Error())
	}
}

func ErrUnauthorized(w http.ResponseWriter, message interface{}) {
	MustEncodeWithStatus(w,
		map[string]*Message{"error": {http.StatusUnauthorized, message}},
		http.StatusUnauthorized)
}

func MustWriteJSONBytes(w http.ResponseWriter, message []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json;charset=utf-8")
	w.Write(message)
}

func MustWriteJSON(w http.ResponseWriter, message string) {
	MustWriteJSONBytes(w, []byte(message))
}
