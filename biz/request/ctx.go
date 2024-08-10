package request

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jhue/misgo/biz/model/base"
)

type BizContext struct {
	a *app.RequestContext
}

func NewBizContext(a *app.RequestContext) BizContext {
	return BizContext{a}
}

func (c *BizContext) Error(code Code, err error) {
	c.Fail(code, err.Error())
}

func (c *BizContext) ParmaError(err error) {
	c.Error(ParmaError, err)
}

func (c *BizContext) DBError(err error) {
	c.Error(DataBaseError, err)
}

func (c *BizContext) ServerError(err error) {
	c.Error(ServerError, err)
}

func (c *BizContext) Fail(code Code, errorMsg string) {
	var msg string
	switch code {
	case Success:
		return
	case ParmaError:
		msg = parmaError.joinMsg(errorMsg)
	case DataBaseError:
		msg = dataBaseError.joinMsg(errorMsg)
	case ServerError:
		msg = serverError.joinMsg(errorMsg)
	}
	errResp := base.CommonResp{
		Code:    int64(code),
		Message: msg,
	}
	c.a.JSON(consts.StatusBadRequest, &errResp)
	c.a.Abort()
}

func (c *BizContext) Text(str string) {
	c.a.String(consts.StatusOK, str)
}

func (c *BizContext) Success() {
	successResp := base.CommonResp{
		Code:    int64(Success),
		Message: "ok",
	}
	c.a.JSON(consts.StatusOK, &successResp)
}

func (c *BizContext) SuccessWithMsg(successMsg string) {
	successResp := base.CommonResp{
		Code:    int64(Success),
		Message: successMsg,
	}
	c.a.JSON(consts.StatusOK, &successResp)
}

func (c *BizContext) Response(obj any) {
	c.a.JSON(consts.StatusOK, obj)
}
