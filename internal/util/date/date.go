package date

import "time"

func Day(src time.Time) time.Time {
	d := time.Date(src.Year(), src.Month(), src.Day(), 0, 0, 0, 0, src.Location())
	return d
}

func Week(src time.Time) time.Time {
	// 获取星期几
	weekday := int(src.Weekday())
	// 计算出当前日期属于该星期的哪一天
	if weekday == 0 {
		weekday = 7
	}
	// 减去星期几的天数，获取星期的第一天（星期一）
	d := time.Date(src.Year(), src.Month(), src.Day()-weekday+1, 0, 0, 0, 0, src.Location())
	return d
}

func Month(src time.Time) time.Time {
	d := time.Date(src.Year(), src.Month(), 1, 0, 0, 0, 0, src.Location())
	return d
}
