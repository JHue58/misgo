// Code generated by hertz generator.

package ping

import (
	"context"
	"github.com/jhue/misgo/internal/mislog"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/jhue/misgo/biz/model/ping"
)

// Ping .
// @router /ping [GET]
func Ping(ctx context.Context, c *app.RequestContext) {
	resp := &ping.PingResp{
		Message: "pong",
		Date:    time.Now().String(),
	}
	c.JSON(consts.StatusOK, resp)
	mislog.DefaultLogger.Info("Ping")
}
