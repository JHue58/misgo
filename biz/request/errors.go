package request

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jhue/misgo/biz/model/base"
)

type ErrorCode struct {
	code   int64
	prefix string
}

func (c ErrorCode) joinMsg(msg string) string {
	return fmt.Sprintf(c.prefix, msg)
}

var (
	ParmaError    = ErrorCode{code: 1, prefix: "参数错误: %s"}
	DataBaseError = ErrorCode{code: 2, prefix: "数据库错误: %s"}
	ServerError   = ErrorCode{code: 3, prefix: "服务器错误: %s"}
)

func ErrorRequest(c *app.RequestContext, code ErrorCode, errorMsg string) {
	errResp := base.ErrorResp{
		Code:    code.code,
		Message: code.joinMsg(errorMsg),
	}
	c.JSON(consts.StatusBadRequest, &errResp)
	c.Abort()
}
