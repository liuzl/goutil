package goutil

import (
	"crypto/md5"
	"encoding/hex"
)

func ContentMD5(value []byte) string {
	hasher := md5.New()
	hasher.Write(value)
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5(text string) string {
	return ContentMD5([]byte(text))
}
