package goutil

import "time"

var (
	TimeZodiac = []string{"水瓶", "双鱼", "白羊", "金牛", "双子", "巨蟹", "狮子", "处女", "天秤", "天蝎", "射手", "魔羯"}
	ZodiacDays = []int{21, 20, 21, 20, 21, 22, 23, 23, 23, 24, 22, 21}

	TimeCnZodiacEmoji = []string{"🐀", "🐂", "🐅", "🐇", "🐉", "🐍", "🐎", "🐐", "🐒", "🐓", "🐕", "🐖"}
	TimeCnZodiac      = []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	TimeGan           = []string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}
	TimeZhi           = []string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}
)

func TimeStr(t int64) string {
	return time.Unix(t, 0).Format("20060102030405")
}

func Sleep(d time.Duration, interrupt chan bool) {
	select {
	case <-interrupt:
		return
	case <-time.After(d):
		return
	}
}

func GetZodiacForTime(t time.Time) string {
	month := t.Month()
	day := t.Day()
	if day < ZodiacDays[month-1] {
		month--
	}
	if month > 0 {
		return TimeZodiac[month-1]
	}
	return TimeZodiac[11]
}

func GetCnZodiacEmojiForTime(t time.Time) string {
	return TimeCnZodiacEmoji[(t.Year()+8)%12]
}

func GetCnZodiacForTime(t time.Time) string {
	return TimeCnZodiac[(t.Year()+8)%12]
}

func GetGanForTime(t time.Time) string {
	return TimeGan[(t.Year()+8)%12]
}

func GetZhiForTime(t time.Time) string {
	return TimeZhi[(t.Year()+6)%10]
}
