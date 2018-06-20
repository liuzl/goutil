package goutil

func ReverseOrderEncode(s string) string {
	var buf []rune
	for _, c := range s {
		if c >= '0' && c <= '9' {
			buf = append(buf, '9'-c+'0')
		} else {
			buf = append(buf, c)
		}
	}
	return string(buf)
}
