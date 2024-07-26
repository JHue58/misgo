package clipboard

import (
	"gorm.io/gorm"
	"time"
)

type ClipBoard struct {
	gorm.Model
	UserID  uint
	Content string `gorm:"not null"`
	Time    time.Time
}
