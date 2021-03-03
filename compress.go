package goutil

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
)

// Base64zlibEncode zlib compress the input string and then base64 encode it
func Base64zlibEncode(content string) string {
	src := []byte(content)
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return base64.StdEncoding.EncodeToString(in.Bytes())
}

// Base64zlibEncode base64 decode the string and zlib decompress it
func Base64zlibDecode(s string) string {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	var out bytes.Buffer
	r, _ := zlib.NewReader(bytes.NewReader(data))
	io.Copy(&out, r)
	return out.String()
}
