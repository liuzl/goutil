package goutil

import (
	"testing"
	"time"
)

func TestZodiacGanZhi(t *testing.T) {
	cases := []struct{ ts, zodiac, gan, zhi, emoji string }{
		{"2018", "狗", "戌", "戊", "🐕"},
		{"1983", "猪", "亥", "癸", "🐖"},
		{"2015", "羊", "未", "乙", "🐐"},
	}

	for _, c := range cases {
		tm, err := time.Parse("2006", c.ts)
		if err != nil {
			t.Error(err)
		}
		if GetZodiacForTime(tm) != c.zodiac {
			t.Errorf("GetZodiacForTime(t):%s != expected:%s",
				GetZodiacForTime(tm), c.zodiac)
		}
		if GetGanForTime(tm) != c.gan {
			t.Errorf("GetGanForTime(t):%s != expected:%s",
				GetGanForTime(tm), c.gan)
		}
		if GetZhiForTime(tm) != c.zhi {
			t.Errorf("GetZhiForTime(t):%s != expected:%s",
				GetZhiForTime(tm), c.zhi)
		}
		if GetZodiacEmojiForTime(tm) != c.emoji {
			t.Errorf("GetZodiacEmojiForTime(t):%s != expected:%s",
				GetZodiacEmojiForTime(tm), c.emoji)
		}
	}
}
