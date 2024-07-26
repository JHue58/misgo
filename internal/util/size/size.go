package size

import "fmt"

const B = 1

const (
	KB = 1000 * B
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
)

const (
	KiB = 1024 * B
	MiB = 1024 * KiB
	GiB = 1024 * MiB
	TiB = 1024 * GiB
	PiB = 1024 * TiB
)

func Format(size int) string {
	var formatStr string
	var value float64

	switch {
	case size >= PB:
		value = float64(size) / float64(PB)
		formatStr = fmt.Sprintf("%.2f PB", value)
	case size >= TB:
		value = float64(size) / float64(TB)
		formatStr = fmt.Sprintf("%.2f TB", value)
	case size >= GB:
		value = float64(size) / float64(GB)
		formatStr = fmt.Sprintf("%.2f GB", value)
	case size >= MB:
		value = float64(size) / float64(MB)
		formatStr = fmt.Sprintf("%.2f MB", value)
	case size >= KB:
		value = float64(size) / float64(KB)
		formatStr = fmt.Sprintf("%.2f KB", value)
	default:
		value = float64(size)
		formatStr = fmt.Sprintf("%.2f B", value)
	}

	return formatStr
}
