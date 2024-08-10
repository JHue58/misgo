package mislog

import (
	"fmt"
	"github.com/jhue/misgo/internal/clock"
	"io"
	"os"
)

func LogFileIO(logPath string, version string) (io.WriteCloser, error) {
	path := fmt.Sprintf("%s/misgo_%s_%s.log", logPath, version, clock.BootTime().Format("2006-01-02-15-04-05"))
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return f, err
}
