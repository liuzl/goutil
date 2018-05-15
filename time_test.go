package goutil

import (
	"testing"
	"time"
)

func TestZodiac(t *testing.T) {
	cases := []struct{ ts, zodiac string }{
		{"20180924", "天秤"},
		{"19830601", "双子"},
		{"19840304", "双鱼"},
	}

	for _, c := range cases {
		tm, err := time.Parse("20060102", c.ts)
		if err != nil {
			t.Error(err)
		}
		if GetZodiacForTime(tm) != c.zodiac {
			t.Errorf("GetZodiacForTime(t):%s != expected:%s",
				GetZodiacForTime(tm), c.zodiac)
		}
	}
}

func TestCnZodiacGanZhi(t *testing.T) {
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
		if GetCnZodiacForTime(tm) != c.zodiac {
			t.Errorf("GetCnZodiacForTime(t):%s != expected:%s",
				GetCnZodiacForTime(tm), c.zodiac)
		}
		if GetGanForTime(tm) != c.gan {
			t.Errorf("GetGanForTime(t):%s != expected:%s",
				GetGanForTime(tm), c.gan)
		}
		if GetZhiForTime(tm) != c.zhi {
			t.Errorf("GetZhiForTime(t):%s != expected:%s",
				GetZhiForTime(tm), c.zhi)
		}
		if GetCnZodiacEmojiForTime(tm) != c.emoji {
			t.Errorf("GetCnZodiacEmojiForTime(t):%s != expected:%s",
				GetCnZodiacEmojiForTime(tm), c.emoji)
		}
	}
}
