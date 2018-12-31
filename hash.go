package goutil

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"

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

func ContentBase62MD5(value []byte) string {
	hasher := md5.New()
	hasher.Write(value)
	b62, _ := basex.NewEncoding(b62alphabet)
	return b62.Encode(hasher.Sum(nil))
}

func Base62MD5(text string) string {
	return ContentBase62MD5([]byte(text))
}

func CrcUint32(text string) uint32 {
	return crc32.ChecksumIEEE([]byte(text))
}

func CRC32(text string) string {
	n := CrcUint32(text)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, n)
	b62, _ := basex.NewEncoding(b62alphabet)
	return b62.Encode(bs)
}
