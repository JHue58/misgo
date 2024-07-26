package monitor

import (
	"fmt"
	"github.com/jhue/misgo/internal/mislog"
	"time"
)

type monitor struct {
	fs     []UpdateFunc
	ticker *time.Ticker
}

func NewMonitor(updateFs ...UpdateFunc) Monitor {
	return &monitor{
		fs:     updateFs,
		ticker: time.NewTicker(time.Second * 5),
	}
}

func (m *monitor) Start() error {
	for _, f := range m.fs {
		ok, msg, err := f()
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("monitor 启动失败 %s", msg)
		}
	}
	go m.monitorLoop()
	return nil

}

func (m *monitor) monitorLoop() {
	for range m.ticker.C {
		for _, f := range m.fs {
			ok, msg, err := f()
			if err != nil {
				mislog.DefaultLogger.Errorf("monitor 执行更新函数失败: %s", err.Error())
			}
			if !ok {
				continue
			}
			mislog.DefaultLogger.Noticef("monitor 监听到更新: %s", msg)
		}
	}
}

func (m *monitor) Stop() {
	m.ticker.Stop()
}
