package goutil

import (
	"fmt"
	"regexp"
	"sync"
)

type Regexps struct {
	sync.Mutex
	items map[string]*regexp.Regexp
}

func (r *Regexps) Compile(pattern string) (*regexp.Regexp, error) {
	r.Lock()
	defer r.Unlock()
	if r.items[pattern] != nil {
		return r.items[pattern], nil
	}
	re, err := regexp.Compile(pattern)
	if err == nil {
		r.items[pattern] = re
	}
	return re, err
}

var pool = &Regexps{items: make(map[string]*regexp.Regexp)}

func RegexpParse(content, pattern string) ([]string, error) {
	re, err := pool.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("re:[%s] error:%+v", pattern, err)
	}
	var ret []string
	res := re.FindAllStringSubmatch(content, -1)
	for i, _ := range res {
		switch {
		case len(res[i]) == 1:
			ret = append(ret, res[i][0])
		case len(res[i]) > 1:
			ret = append(ret, res[i][1:]...)
		}
	}
	return ret, nil
}

func RegexpMatch(content, pattern string) bool {
	re, err := pool.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(content)
}
