package conf

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v2"
	"os"
	"reflect"
	"strings"
)

var conf *Config

func InitConfig(v ...string) error {
	c := initConfig()

	if len(v) > 0 {
		c.Version = v[0]
	} else {
		c.Version = "test"
	}

	conf = &c
	if isEmpty(reflect.ValueOf(c)) {
		return fmt.Errorf("config 存在字段为定义")
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
	Level   string `yaml:"level"`
	LogPath string `yaml:"logPath"`
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
	// 读取 YAML 文件
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		data, err = os.ReadFile("./volume/config.yaml")
		if err != nil {
			panic(err)
		}
	}

	// 解析 YAML 数据
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func GetConfig() Config {
	return *conf
}

func isEmpty(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isEmpty(v.Field(i)) {
				return true
			}
		}
		return false
	default:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	}
}
