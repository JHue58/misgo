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
	RecordConfig    `yaml:"recordConfig"`
	MoneyConfig     `yaml:"moneyConfig"`
}

type ClipBoardConfig struct {
	MaxStore      int64  `yaml:"maxStore"`
	FileStorePath string `yaml:"fileStorePath"`
	MaxFileSizeMb int64  `yaml:"maxFileSizeMb"`
}

type RecordConfig struct {
	MaxGetCount int64 `yaml:"maxGetCount"`
}

type MoneyConfig struct {
	MaxGetCount int64 `yaml:"maxGetCount"`
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
