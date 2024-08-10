package clock

import "time"

var bootTime = time.Now()

func BootSince() time.Duration {
	return time.Since(bootTime)
}

func BootTime() time.Time {
	return bootTime
}
