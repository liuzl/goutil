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

func (this *ResponseProxyWriter) Header() http.Header {
	return this.writer.Header()
}

func (this *ResponseProxyWriter) Write(bytes []byte) (int, error) {
	if !this.wroteHeader {
		this.WriteHeader(http.StatusOK)
	}
	this.Body = append(this.Body, bytes[0:len(bytes)]...)
	return this.writer.Write(bytes)
}

func (this *ResponseProxyWriter) WriteHeader(i int) {
	if this.wroteHeader {
		return
	}
	this.wroteHeader = true

	this.Code = i
	this.writer.WriteHeader(i)
}

func NewRespProxyWriter(w http.ResponseWriter) *ResponseProxyWriter {
	return &ResponseProxyWriter{
		writer: w,
		Body:   []byte{},
	}
}
