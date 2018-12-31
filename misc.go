package goutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

// ReverseOrderEncode encodes digits string
func ReverseOrderEncode(s string) string {
	var buf []rune
	for _, c := range s {
		if c >= '0' && c <= '9' {
			buf = append(buf, '9'-c+'0')
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}

// FileGuard is used for file lock
func FileGuard(f string) bool {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return false
	}
	fname := filepath.Join(dir, f)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		if err = ioutil.WriteFile(fname,
			[]byte(time.Now().UTC().Format(time.RFC3339)), 0644); err != nil {
			return false
		}
		return true
	}
	return false
}
