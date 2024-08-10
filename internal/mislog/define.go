package mislog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jhue/misgo/pkg/formater"
)

type Logger interface {
	hlog.Logger
	hlog.FormatLogger
	Snapshot(item formater.Item) error
}

type Level hlog.Level
