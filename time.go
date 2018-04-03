package goutil

import "time"

func Sleep(d time.Duration, interrupt chan bool) {
	select {
	case <-interrupt:
		return
	case <-time.After(d):
		return
	}
}
