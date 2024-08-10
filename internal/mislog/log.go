package mislog

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jhue/misgo/internal/conf"
	"github.com/jhue/misgo/pkg/formater"
	"io"
	"os"
	"time"
)

var DefaultLogger Logger

type logger struct {
	stdLog hlog.FullLogger
	level  hlog.Level
	c      *controller
	f      io.Writer
}

func InitLogger(f io.Writer, config conf.LogConfig) {
	fileWriter := io.MultiWriter(f, os.Stdout)
	hlog.SetLevel(config.HLevel())
	hlog.SetOutput(fileWriter)
	DefaultLogger = &logger{
		stdLog: hlog.DefaultLogger(),
		level:  config.HLevel(),
		c:      newController(10, hlog.LevelInfo, time.Second),
		f:      f,
	}

}

func (l *logger) Trace(v ...interface{}) {
	if !l.c.Allow(hlog.LevelTrace) {
		return
	}
	l.stdLog.Trace(v...)
}

func (l *logger) Debug(v ...interface{}) {
	if !l.c.Allow(hlog.LevelDebug) {
		return
	}
	l.stdLog.Debug(v...)
}

func (l *logger) Info(v ...interface{}) {
	if !l.c.Allow(hlog.LevelInfo) {
		return
	}
	l.stdLog.Info(v...)
}

func (l *logger) Notice(v ...interface{}) {
	if !l.c.Allow(hlog.LevelNotice) {
		return
	}
	l.stdLog.Notice(v...)
}

func (l *logger) Warn(v ...interface{}) {
	if !l.c.Allow(hlog.LevelWarn) {
		return
	}
	l.stdLog.Warn(v...)
}

func (l *logger) Error(v ...interface{}) {
	if !l.c.Allow(hlog.LevelError) {
		return
	}
	l.stdLog.Error(v...)
}

func (l *logger) Fatal(v ...interface{}) {
	if !l.c.Allow(hlog.LevelFatal) {
		return
	}
	l.stdLog.Fatal(v...)
}

func (l *logger) Tracef(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelTrace) {
		return
	}
	l.stdLog.Tracef(format, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelDebug) {
		return
	}
	l.stdLog.Debugf(format, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelInfo) {
		return
	}
	l.stdLog.Infof(format, v...)
}

func (l *logger) Noticef(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelNotice) {
		return
	}
	l.stdLog.Noticef(format, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelWarn) {
		return
	}
	l.stdLog.Warnf(format, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelError) {
		return
	}
	l.stdLog.Errorf(format, v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	if !l.c.Allow(hlog.LevelFatal) {
		return
	}
	l.stdLog.Fatalf(format, v...)
}

func (l *logger) Snapshot(item formater.Item) (err error) {
	config := conf.GetConfig()

	lineCount := config.SnapshotLineCount

	file, ok := l.f.(*os.File)
	if !ok {
		return fmt.Errorf("logIO is %T, but not *os.File ", file)
	}
	path := file.Name()
	file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {

		return
	}

	// 文件的大小
	fileSize := stat.Size()

	// 逐行从文件末尾向前读取
	buf := make([]byte, 1)
	bufs := make([]byte, 0)

	var lined int
	for i := fileSize - 1; i >= 0; i-- {
		if lined >= lineCount {
			break
		}
		_, err := file.Seek(i, 0)
		if err != nil {
			return err
		}
		_, err = file.Read(buf)
		if err != nil {
			return err
		}

		if buf[0] == '\n' {
			lined += 1
		}
		bufs = append(bufs, buf[0])
	}
	length := len(bufs)
	r := make([]byte, length)

	idx := 0
	for i := 0; i < length; i++ {
		r[i] = bufs[length-i-1]
		if r[i] == '\n' {
			item.WriteLineBytes(r[idx:i])
			idx = i + 1
			// 跳过换行
		}

	}
	item.Done()
	return nil
}
