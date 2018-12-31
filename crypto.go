package goutil

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/eknkc/basex"
)

var b62alphabet = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func ContentMD5(value []byte) string {
	hasher := md5.New()
	hasher.Write(value)
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5(text string) string {
	return ContentMD5([]byte(text))
}

func Base62MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	b62, _ := basex.NewEncoding(b62alphabet)
	return b62.Encode(hasher.Sum(nil))
}
