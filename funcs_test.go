package goutil

import (
	"testing"
)

func sum(x, y int) int {
	return x + y
}

func product(x, y int) int {
	return x * y
}

func TestCall(t *testing.T) {
	funcs := map[string]interface{}{
		"nf.math.sum":     sum,
		"nf.math.product": product,
	}
	ret, err := Call(funcs, "nf.math.sum", 4, 5)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ret)
	for _, v := range ret {
		t.Log(v)
	}
}
