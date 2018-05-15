package goutil

import "time"

var (
	TimeZodiacEmoji = []string{
		"🐀", "🐂", "🐅", "🐇", "🐉", "🐍", "🐎", "🐐", "🐒", "🐓", "🐕", "🐖"}
	TimeZodiac = []string{
		"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	TimeGan = []string{
		"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	TimeZhi = []string{
		"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
)

func Sleep(d time.Duration, interrupt chan bool) {
	select {
	case <-interrupt:
		return
	case <-time.After(d):
		return
	}
}
func GetZodiacEmojiForTime(t time.Time) string {
	return TimeZodiacEmoji[(t.Year()+8)%12]
}

func GetZodiacForTime(t time.Time) string {
	return TimeZodiac[(t.Year()+8)%12]
}

func GetGanForTime(t time.Time) string {
	return TimeGan[(t.Year()+8)%12]
}

func GetZhiForTime(t time.Time) string {
	return TimeZhi[(t.Year()+6)%10]
}
