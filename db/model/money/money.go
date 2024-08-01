package money

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

var model = Transaction{}

type Transaction struct {
	gorm.Model
	UserID uint `gorm:"index"`
	// 支出 or 收入
	Type string `gorm:"not null"`
	// 用途/来源
	Category string `gorm:"not null"`
	// 金额
	Amount float64 `gorm:"not null"`
	// 备注
	Note string    `gorm:"not null"`
	Time time.Time `gorm:"not null"`
}

func (t *Transaction) String() string {
	formatTime := t.Time.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("%s %s %s %.2f元 %s", formatTime, t.Category, t.Type, t.Amount, t.Note)
}

// Markdown 表格
//
//	|日期|类型（支出/收入）|用途/来源|金额|备注|
func (t *Transaction) Markdown() string {
	formatTime := t.Time.Format("2006-01-02 15:04:05")
	return fmt.Sprintf("| %s | %s | %s | %.2f | %s | \n",
		formatTime, t.Type, t.Category, t.Amount, t.Note)
}

func Model() *Transaction {
	return &model
}
