package money

import (
	"fmt"
	"github.com/jhue/misgo/internal/util"
	"time"
)

func (x *Transaction) Filter() {
	x.Category = util.PopEmojis(x.Category)
}

func (x *TransactionGetReq) GetDateRange() (startDate time.Time, endDate time.Time, format string) {
	now := time.Now()

	switch x.TimeRange {
	case "今天":
		startDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		endDate = now
		format = "今天"
	case "昨天":
		yesterday := now.AddDate(0, 0, -1)
		startDate = time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, now.Location())
		endDate = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, -1, now.Location())
		format = "昨天"
	case "这周", "本周":
		startOfWeek := now.AddDate(0, 0, -int(now.Weekday()))
		startDate = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, now.Location())
		endDate = now
		format = "这周"
	case "上周":
		startOfWeek := now.AddDate(0, 0, -int(now.Weekday())-7)
		endOfWeek := startOfWeek.AddDate(0, 0, 7)
		startDate = time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, now.Location())
		endDate = time.Date(endOfWeek.Year(), endOfWeek.Month(), endOfWeek.Day(), 0, 0, 0, -1, now.Location())
		format = "上周"
	case "这个月":
		startDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		endDate = now
		format = "这个月"
	case "上个月":
		startOfLastMonth := now.AddDate(0, -1, 0)
		startDate = time.Date(startOfLastMonth.Year(), startOfLastMonth.Month(), 1, 0, 0, 0, 0, now.Location())
		endDate = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, -1, now.Location())
		format = "上个月"
	default:
		// 使用传入的时间戳作为时间范围
		startDate = time.Unix(x.StartTime, 0)
		endDate = time.Unix(x.EndTime, 0)
		format = fmt.Sprintf("%s 到 %s", startDate.Format("2006年01月02日"), endDate.Format("2006年01月02日"))
	}

	return
}
