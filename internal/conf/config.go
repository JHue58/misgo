package conf

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/jhue/misgo/internal/util"
	"strings"
	"sync/atomic"
)

var conf atomic.Value

var paths = []string{
	"./config.yaml",
	"./volume/config.yaml",
}

func InitConfig(v ...string) error {
	c := initConfig()

	if len(v) > 0 {
		c.Version = v[0]
	} else {
		c.Version = "test"
	}
	err := util.EmptyCheck(c)
	if err != nil {
		return err
	}

	if GetConfig() != c {
		conf.Store(&c)
	}
	return nil
}

type Config struct {
	Version      string
	ServerConfig `yaml:"serverConfig"`
	DBConfig     `yaml:"dbConfig"`
	LogConfig    `yaml:"logConfig"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DBConfig struct {
	Path string `yaml:"path"`
	// 最大空闲连接数
	IdleConn int `yaml:"idleConn"`
	// 最大连接数
	MaxConn int `yaml:"maxConn"`
	// 连接最大存活时间(分钟)
	MaxLifeTime int64 `yaml:"maxLifeTime"`
}

type LogConfig struct {
	Level             string `yaml:"level"`
	LogPath           string `yaml:"logPath"`
	SnapshotLineCount int    `yaml:"snapshotLineCount"`
}

func (c LogConfig) HLevel() (level hlog.Level) {
	switch strings.ToLower(c.Level) {
	case "debug":
		level = hlog.LevelDebug
	case "info":
		level = hlog.LevelInfo
	case "warn":
		level = hlog.LevelWarn
	case "error":
		level = hlog.LevelError
	case "fatal":
		level = hlog.LevelFatal
	case "notice":
		level = hlog.LevelNotice
	case "trace":
		level = hlog.LevelTrace
	default:
		level = hlog.LevelInfo
	}
	return
}

func initConfig() Config {

	// 解析 YAML 数据
	var config Config
	err := util.ReadYaml(&config, paths...)
	if err != nil {
		panic(err)
	}

	return config
}

func GetConfig() Config {
	c, ok := conf.Load().(*Config)
	if !ok {
		return Config{}
	} else {
		return *c
	}
}
