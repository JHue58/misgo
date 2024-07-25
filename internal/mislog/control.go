package mislog

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"sync/atomic"
	"time"
)

type controller struct {
	nowPass     atomic.Int64
	canPass     int64
	rejectLevel hlog.Level
	ticker      *time.Ticker
}

func newController(pass int64, rejectLevel hlog.Level, interval time.Duration) *controller {
	c := &controller{rejectLevel: rejectLevel, ticker: time.NewTicker(interval), canPass: pass}
	go c.tickerLoop()
	return c
}

func (c *controller) tickerLoop() {
	for range c.ticker.C {
		c.nowPass.Store(0)
	}
}

func (c *controller) Allow(level hlog.Level) bool {
	added := c.nowPass.Add(1)
	if level > c.rejectLevel {
		return true
	}
	if added > c.canPass {
		return false
	}
	return true
}

func (c *controller) Close() error {
	c.ticker.Stop()
	return nil
}
