package mislog

import "github.com/cloudwego/hertz/pkg/common/hlog"

type Logger interface {
	hlog.Logger
	hlog.FormatLogger
	Snapshot() (string, error)
}

type Level hlog.Level
