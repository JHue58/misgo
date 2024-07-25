package mislog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jhue/misgo/internal/conf"
	"io"
	"os"
	"time"
)

var DefaultLogger Logger

type logger struct {
	stdLog hlog.FullLogger
	level  hlog.Level
	c      *controller
}

func InitLogger(f io.Writer, config conf.LogConfig) {

	fileWriter := io.MultiWriter(f, os.Stdout)
	hlog.SetLevel(config.HLevel())
	hlog.SetOutput(fileWriter)
	DefaultLogger = &logger{
		stdLog: hlog.DefaultLogger(),
		level:  config.HLevel(),
		c:      newController(10, hlog.LevelInfo, time.Second),
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
