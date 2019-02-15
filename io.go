package goutil

import (
	"bufio"
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

// ForEachLine higher order function that processes each line of text by callback function.
// The last non-empty line of input will be processed even if it has no newline.
func ForEachLine(br *bufio.Reader, callback func(string) error) error {
	stop := false
	for {
		if stop {
			break
		}
		line, err := br.ReadString('\n')
		if err == io.EOF {
			stop = true
		} else if err != nil {
			return err
		}
		line = strings.TrimSuffix(line, "\n")
		if line == "" {
			if !stop {
				if err = callback(line); err != nil {
					return err
				}
			}
			continue
		}
		if err = callback(line); err != nil {
			return err
		}
	}
	return nil
}
