package goutil

import "time"

var (
	// TimeZodiac stores 12 Astrology Zodiac in Chinese
	TimeZodiac = []string{"水瓶", "双鱼", "白羊", "金牛", "双子", "巨蟹", "狮子", "处女", "天秤", "天蝎", "射手", "魔羯"}
	zodiacDays = []int{21, 20, 21, 20, 21, 22, 23, 23, 23, 24, 22, 21}

	// TimeCnZodiacEmoji stores the 12 Chinese Emoji Zodiac
	TimeCnZodiacEmoji = []string{"🐀", "🐂", "🐅", "🐇", "🐉", "🐍", "🐎", "🐐", "🐒", "🐓", "🐕", "🐖"}
	// TimeCnZodiac stores the 12 Chinese Zodiac
	TimeCnZodiac = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	// TimeGan stores the 12 Chinese Gan
	TimeGan = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	// TimeZhi stores the 10 Chinese Zhi
	TimeZhi = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
)

// TimeStr returns the 20060102150405 format time for timestamp t in UTC
func TimeStr(t int64) string {
	return time.Unix(t, 0).UTC().Format("20060102150405")
}

// Sleep is a interruptable version sleep
// if interrupted, return true, otherwise return false
func Sleep(d time.Duration, interrupt chan bool) bool {
	select {
	case <-interrupt:
		return true
	case <-time.After(d):
		return false
	}
}

// GetZodiacForTime returns the Zodiac for time t in Chinese
func GetZodiacForTime(t time.Time) string {
	month := t.Month()
	day := t.Day()
	if day < zodiacDays[month-1] {
		month--
	}
	if month > 0 {
		return TimeZodiac[month-1]
	}
	return TimeZodiac[11]
}

// GetCnZodiacEmojiForTime returns the Chinese Zodiac for time t in Emoji
func GetCnZodiacEmojiForTime(t time.Time) string {
	return TimeCnZodiacEmoji[(t.Year()+8)%12]
}

// GetCnZodiacForTime returns the Chinese Zodiac for time t in Chinese
func GetCnZodiacForTime(t time.Time) string {
	return TimeCnZodiac[(t.Year()+8)%12]
}

// GetGanForTime returns the Gan for time t in Chinese
func GetGanForTime(t time.Time) string {
	return TimeGan[(t.Year()+8)%12]
}

// GetZhiForTime returns the Zhi for time t in Chinese
func GetZhiForTime(t time.Time) string {
	return TimeZhi[(t.Year()+6)%10]
}
