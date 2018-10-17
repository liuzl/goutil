package goutil

import (
	"testing"
)

func TestBase62(t *testing.T) {
	cases := map[int64]string{
		0:             "0",
		12345:         "3d7",
		6285600000447: "1MF0R2gv",
	}
	for k, v := range cases {
		v1 := ToB62(k)
		if v1 != v {
			t.Error(v, v1)
		}
		k1, err := FromB62(v)
		if err != nil {
			t.Error(err)
		}
		if k1 != k {
			t.Error(k, k1)
		}
	}
}
