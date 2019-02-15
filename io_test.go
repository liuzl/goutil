package goutil

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileLineCount(t *testing.T) {
	cases := map[string]int{
		"io.go": 83,
	}
	for k, v := range cases {
		cnt, err := FileLineCount(k)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, v, cnt, "should be equal")
	}
}

func TestForEachLine(t *testing.T) {
	fr, err := os.Open("io.go")
	if err != nil {
		t.Error(err)
	}
	r := bufio.NewReader(fr)
	err = ForEachLine(r, func(line string) error {
		t.Log(line)
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
