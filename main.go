// Code generated by hertz generator.

package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/jhue/misgo/biz"
	"github.com/jhue/misgo/biz/middleware"
	"github.com/jhue/misgo/internal/conf"
	"github.com/jhue/misgo/internal/mislog"
	"github.com/jhue/misgo/pkg/monitor"
)

func main() {
	m := newMonitor()
	err := m.Start()
	if err != nil {
		panic(err)
	}
	defer m.Stop()

	config := conf.GetConfig()

	f, err := mislog.LogFileIO(config.LogPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	mislog.InitLogger(f, config.LogConfig)
	mislog.DefaultLogger.Infof("misgo version: %s", config.Version)
	h := server.Default(server.WithHostPorts(fmt.Sprintf("%s:%d", config.Host, config.Port)))
	h.Use(middleware.UIDExtractMiddleware())
	register(h)
	h.Spin()
}

func newMonitor() monitor.Monitor {
	updateFuncs := []monitor.UpdateFunc{
		func() (ok bool, msg string, err error) {
			before := conf.GetConfig()
			err = conf.InitConfig(Version)
			if err != nil {
				return
			}
			after := conf.GetConfig()
			if before != after {
				return true, "config修改成功", nil
			} else {
				return
			}
		},
		func() (ok bool, msg string, err error) {
			before := biz.GetBizConfig()
			err = biz.InitBizConfig()
			if err != nil {
				return
			}
			after := biz.GetBizConfig()
			if before != after {
				return true, "bizConfig修改成功", nil
			} else {
				return
			}
		},
	}
	return monitor.NewMonitor(updateFuncs...)
}
