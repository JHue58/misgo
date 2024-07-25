package mislog

import "github.com/cloudwego/hertz/pkg/common/hlog"

type Logger interface {
	hlog.Logger
	hlog.FormatLogger
}

type Level hlog.Level
