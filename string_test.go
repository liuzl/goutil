package goutil

import (
	"testing"
)

func TestJoin(t *testing.T) {
	cases := [][]string{
		[]string{},
		[]string{"hello", "world"},
		[]string{"天", "津", "大", "学"},
	}
	expects := []string{"", "hello world", "天津大学"}

	for i := 0; i < len(cases); i++ {
		if Join(cases[i]) != expects[i] {
			t.Error(expects[i], cases[i])
		}
	}
}
