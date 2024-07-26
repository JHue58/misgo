package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/jhue/misgo/biz/request"
	"github.com/jhue/misgo/db"
	"github.com/jhue/misgo/db/model/user"
)

var emptyUID = uid{}

type uid struct {
	UID string `json:"uid"`
}

func (u uid) IsEmpty() bool {
	return u == emptyUID
}

func UIDExtractMiddleware() func(ctx context.Context, c *app.RequestContext) {
	d := db.Get()
	return func(ctx context.Context, c *app.RequestContext) {
		biz := request.NewBizContext(c)
		var u uid
		err := c.BindAndValidate(&u)
		if err != nil || u.IsEmpty() {
			return
		}

		usr := user.User{}
		res := d.Find(&usr, "uid = ?", u.UID)
		if res.Error != nil || res.RowsAffected == 0 {
			biz.Fail(request.ParmaError, fmt.Sprintf("uid '%s' 不存在", u.UID))

			return
		}
		c.Next(user.WithUser(ctx, usr))

	}
}
