package rest

import (
	"encoding/json"
	"net/http"
)

type RestMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
