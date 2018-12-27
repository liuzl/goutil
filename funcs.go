package goutil

import (
	"fmt"
	"reflect"
)

func Call(m map[string]interface{},
	name string, params ...interface{}) ([]reflect.Value, error) {
	var nf interface{}
	if nf = m[name]; nf == nil {
		return nil, fmt.Errorf("func %s not found", name)
	}
	f := reflect.ValueOf(nf)
	if f.Kind() != reflect.Func {
		return nil, fmt.Errorf("%s is not a function", name)
	}
	t := f.Type()
	if !t.IsVariadic() && len(params) != t.NumIn() {
		return nil, fmt.Errorf("(len(params)=%d) != (t.NumIn()=%d)",
			len(params), t.NumIn())
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
		if !t.IsVariadic() && in[k].Type() != t.In(k) {
			return nil, fmt.Errorf("(in[%d].Type()=%s) != (t.In(%d)=%s)",
				k, in[k].Type(), k, t.In(k))
		}
	}
	return f.Call(in), nil
}
