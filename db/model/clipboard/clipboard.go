package clipboard

import (
	"gorm.io/gorm"
	"time"
)

const (
	TextType int = iota
	FileType
)

type ClipBoard struct {
	gorm.Model
	UserID  uint
	Type    int    // 类型，是文件类型的时候Content为文件名
	Content string `gorm:"not null"`
	Hash    string
	Time    time.Time
}
