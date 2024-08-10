package mislog

import (
	"fmt"
	"io"
	"os"
	"time"
)

var timeStr = time.Now().Format("2006-01-02-15-04-05")

func LogFileIO(logPath string, version string) (io.WriteCloser, error) {
	path := fmt.Sprintf("%s/misgo_%s_%s.log", logPath, version, timeStr)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return f, err
}
