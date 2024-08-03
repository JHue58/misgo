package date

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	now := time.Now()
	time.Sleep(time.Second)
	after := time.Now()
	fmt.Println(Day(now))
	fmt.Println(Day(after))
	fmt.Println(Day(now) == Day(after))
	mp := make(map[time.Time]int)
	mp[Day(now)] = 1
	mp[Day(time.Now())] = 2
	fmt.Println(mp[Day(time.Now())])
}
