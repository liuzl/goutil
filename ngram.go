package goutil

import (
	"github.com/liuzl/ling"
	"strings"
)

var nlp = ling.MustNLP(ling.Norm)

func GetNGramFromArray(min, max int, words []string) map[string]int {
	if min <= 0 || max <= 0 || min > max {
		return nil
	}
	dict := make(map[string]int)
	n := len(words)
	for i := 0; i < n; i++ {
		for j := min; j <= max; j++ {
			if i+j > n {
				break
			}
			dict[strings.Join(words[i:i+j], " ")]++
		}
	}
	return dict
}

func GetNGram(min, max int, input string) map[string]int {
	d := ling.NewDocument(input)
	if err := nlp.Annotate(d); err != nil {
		return nil
	}
	return GetNGramFromArray(min, max, d.XRealTokens(ling.Norm))
}
