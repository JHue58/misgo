package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sync/atomic"
)

func FlowControlMiddleware(threshold int64) func(ctx context.Context, c *app.RequestContext) {
	var inProcess atomic.Int64
	return func(ctx context.Context, c *app.RequestContext) {
		if inProcess.Load() > threshold {
			c.AbortWithMsg("service is busy.", consts.StatusTooManyRequests)
			return
		}
		inProcess.Add(1)
		defer inProcess.Add(-1)
		c.Next(ctx)
	}
}
