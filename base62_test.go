package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase62(t *testing.T) {
	cases := map[uint64]string{
		0:             "0",
		12345:         "3d7",
		6285600000447: "1MF0R2gv",
	}
	for k, v := range cases {
		v1 := ToB62(k)
		assert.Equal(t, v, v1, "should be equal")
		k1, err := FromB62(v)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, k, k1, "should be equal")
	}
}
