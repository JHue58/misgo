// Code generated by hertz generator.

package record

import (
	"context"
	"github.com/jhue/misgo/biz/request"
	"github.com/jhue/misgo/db"
	recordM "github.com/jhue/misgo/db/model/record"
	"github.com/jhue/misgo/db/model/user"
	"github.com/jhue/misgo/internal/mislog"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jhue/misgo/biz/model/record"
)

// Record .
// @router /record [PUT]
func Record(ctx context.Context, c *app.RequestContext) {
	var err error
	var req record.RecordReq
	err = c.BindAndValidate(&req)
	if err != nil {
		request.ErrorRequest(c, request.ParmaError, err.Error())
		return
	}
	user, ok := user.ExtractUser(ctx)
	if !ok {
		request.ErrorRequest(c, request.ParmaError, "key不正确或不存在")
		return
	}
	if req.Tag == "" || req.Content == "" {
		request.ErrorRequest(c, request.ParmaError, "tag和content不能为空")
		return
	}

	d := db.Get()
	res := d.Create(&recordM.Record{
		UserID:  user.ID,
		Tag:     req.Tag,
		Content: req.Content,
		Extend:  req.Extend,
		Time:    time.Now(),
	})

	if res.Error != nil {
		request.ErrorRequest(c, request.DataBaseError, res.Error.Error())
		return
	}

	resp := &record.RecordResp{Content: "ok"}

	c.JSON(consts.StatusOK, resp)

	mislog.DefaultLogger.Infof("Record Success [Name] %s [Tag] %s [Content] %s\n", user.Name, req.Tag, req.Content)
}
