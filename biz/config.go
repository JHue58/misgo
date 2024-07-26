package biz

import (
	"github.com/jhue/misgo/internal/util"
	"sync/atomic"
)

var conf atomic.Value

var paths = []string{
	"./config_biz.yaml",
	"./volume/config_biz.yaml",
}

type Config struct {
	ClipBoardConfig `yaml:"clipBoardConfig"`
}

type ClipBoardConfig struct {
	MaxStore int64 `yaml:"maxStore"`
}

func InitBizConfig() error {
	c := initBizConfig()
	err := util.EmptyCheck(c)
	if err != nil {
		return err
	}
	if GetBizConfig() != c {
		conf.Store(&c)
	}
	return nil
}

func initBizConfig() Config {
	// 解析 YAML 数据
	var config Config
	err := util.ReadYaml(&config, paths...)
	if err != nil {
		panic(err)
	}
	return config
}

func GetBizConfig() Config {
	c, ok := conf.Load().(*Config)
	if !ok {
		return Config{}
	} else {
		return *c
	}

}
