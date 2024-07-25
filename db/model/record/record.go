package record

import "time"

type Record struct {
	UserID  uint
	Tag     string `gorm:"not null"`
	Content string `gorm:"not null"`
	Extend  string
	Time    time.Time
}
