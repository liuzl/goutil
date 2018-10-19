package goutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, GetZodiacForTime(tm), c.zodiac, "")
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
		assert.Equal(t, GetCnZodiacForTime(tm), c.zodiac, "")
		assert.Equal(t, GetGanForTime(tm), c.gan, "")
		assert.Equal(t, GetZhiForTime(tm), c.zhi, "")
		assert.Equal(t, GetCnZodiacEmojiForTime(tm), c.emoji, "")
	}
}
