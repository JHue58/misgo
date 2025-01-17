package mislog

import (
	"fmt"
	"io"
	"os"
	"time"
)

func LogFileIO(logPath string) (io.WriteCloser, error) {
	path := fmt.Sprintf("%s/misgo_%s.log", logPath, time.Now().Format("2006-01-02 15-04-05"))
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return f, err
}
