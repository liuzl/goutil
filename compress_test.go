package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64zlib(t *testing.T) {
	cases := []string{"123123123", "abcdef", "123asdfasd", "xyzddad"}
	for _, s := range cases {
		content := ""
		for i := 0; i < 100; i++ {
			content += s
		}
		assert.Equal(t, content, Base64zlibDecode(Base64zlibEncode(content)), "should be equal")
	}
}
