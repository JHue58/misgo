package clipboard

import "time"

type ClipBoard struct {
	UserID  uint
	Content string `gorm:"not null"`
	Time    time.Time
}
