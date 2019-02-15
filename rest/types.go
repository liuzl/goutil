package rest

import (
	"fmt"
	"net/http"
)

type ResponseLog struct {
	Request    *http.Request
	StatusCode int    `json:"staus_code"`
	Body       string `json:"body"`
	Header     string `json:"header"`
}

func (rl ResponseLog) DumpResponse() string {
	res := ""
	res += fmt.Sprintf("HTTP/%d.%d %d %s\r\n", rl.Request.ProtoMajor, rl.Request.ProtoMinor,
		rl.StatusCode, http.StatusText(rl.StatusCode))
	res += rl.Header
	res += "\r\n"
	res += rl.Body
	return res
}

type ResponseProxyWriter struct {
	writer      http.ResponseWriter
	Body        []byte
	Code        int
	wroteHeader bool
}

func (w *ResponseProxyWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *ResponseProxyWriter) Write(bytes []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}
	w.Body = append(w.Body, bytes[0:len(bytes)]...)
	return w.writer.Write(bytes)
}

func (w *ResponseProxyWriter) WriteHeader(i int) {
	if w.wroteHeader {
		return
	}
	w.wroteHeader = true

	w.Code = i
	w.writer.WriteHeader(i)
}

func NewRespProxyWriter(w http.ResponseWriter) *ResponseProxyWriter {
	return &ResponseProxyWriter{
		writer: w,
		Body:   []byte{},
	}
}
