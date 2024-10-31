package goutil

import (
	"crypto/rand"
	r "math/rand"
	"strings"
	"time"
	"unicode"
)

type is func(rune) bool

// StringIs checkes whether each runes of s satisfy f
func StringIs(s string, f is) bool {
	for _, c := range s {
		if !f(c) {
			return false
		}
	}
	return true
}

// IsASCII checkes whether r is within ASCII or not
func IsASCII(r rune) bool {
	return r <= unicode.MaxASCII
}

// IsLatin1 checkes whether r is within Latin1 or not
func IsLatin1(r rune) bool {
	return r <= unicode.MaxLatin1
}

// IsCJK checkes whether r is a CJK rune or not
func IsCJK(r rune) bool {
	return unicode.Is(unicode.Scripts["Han"], r)
}

var noSpaceScripts = []string{"Han", "Lao", "Thai", "Tibetan"}

// NoSpaceWriting checkes whether r is within Han, Lao, Thai and Tibetan
func NoSpaceWriting(r rune) bool {
	if unicode.IsPunct(r) {
		return true
	}
	for _, s := range noSpaceScripts {
		if unicode.Is(unicode.Scripts[s], r) {
			return true
		}
	}
	return false
}

// Join concates s smartly
func Join(s []string) string {
	if len(s) == 0 {
		return ""
	}
	var ss []string
	pre := false
	for _, str := range s {
		cur := StringIs(str, func(r rune) bool { return !NoSpaceWriting(r) })
		if pre && cur {
			ss = append(ss, " ")
		}
		ss = append(ss, str)
		pre = cur
	}
	return strings.Join(ss, "")
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// GenerateRandomString generate random strings of given length n
func GenerateRandomString(n int) string {
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	letterLen := len(letterBytes)
	for i, b := range bytes {
		if randBy {
			bytes[i] = letterBytes[r.Intn(letterLen)]
		} else {
			bytes[i] = letterBytes[b%byte(letterLen)]
		}
	}
	return string(bytes)
}
