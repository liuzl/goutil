package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	cases := [][]string{
		[]string{},
		[]string{"hello", "world"},
		[]string{"天", "津", "大", "学"},
	}
	expects := []string{"", "hello world", "天津大学"}

	for i := 0; i < len(cases); i++ {
		assert.Equal(t, Join(cases[i]), expects[i], "")
	}
}

func TestRandString(t *testing.T) {
	t.Log(GenerateRandomString(20))
	t.Log(GenerateRandomString(40))
}
