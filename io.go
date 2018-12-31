package goutil

import (
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"strings"
)

// LineCount counts the number of '\n' for r
func LineCount(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// FileLineCount counts the number of '\n' for file f
// f could be gzip file or plain text file
func FileLineCount(f string) (int, error) {
	if strings.HasSuffix(strings.ToLower(f), ".gz") {
		fr, err := os.Open(f)
		if err != nil {
			return 0, err
		}
		defer fr.Close()
		r, err := gzip.NewReader(fr)
		if err != nil {
			return 0, err
		}
		return LineCount(r)
	}
	r, err := os.Open(f)
	if err != nil {
		return 0, err
	}
	defer r.Close()
	return LineCount(r)
}
