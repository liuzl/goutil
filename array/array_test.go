package array

import (
	"testing"
)

func TestArrayIntersect(t *testing.T) {
	var arr []string
	ret, ok := Intersect(arr)
	if !ok {
		t.Error("not ok")
	} else {
		t.Log(ret)
	}
}
