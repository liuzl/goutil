package goutil

import (
	"strings"
	"unicode"
)

type is func(rune) bool

func StringIs(s string, f is) bool {
	for _, c := range s {
		if !f(c) {
			return false
		}
	}
	return true
}

func IsASCII(r rune) bool {
	return r <= unicode.MaxASCII
}

func IsLatin1(r rune) bool {
	return r <= unicode.MaxLatin1
}

func IsCJK(r rune) bool {
	if unicode.Is(unicode.Scripts["Han"], r) {
		return true
	}
	return false
}

var noSpaceScripts = []string{"Han", "Lao", "Thai", "Tibetan"}

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
