package db

import (
	"fmt"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {

	now := time.Now()
	todaySixAM := time.Date(now.Year(), now.Month(), now.Day(), -1, 0, 0, 0, now.Location())
	fmt.Println(todaySixAM.Unix())

}
