package biz

import (
	"github.com/jhue/misgo/internal/util"
	"slices"
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

func (c Config) Equal(target Config) (equal bool) {
	if c.ClipBoardConfig != target.ClipBoardConfig {
		return false
	}
	if c.RecordConfig != target.RecordConfig {
		return false
	}
	if c.MoneyConfig.MaxGetCount != target.MoneyConfig.MaxGetCount {
		return false
	}

	srcParse, dstParse := c.MoneyConfig.ParserConfig, target.MoneyConfig.ParserConfig

	if !slices.Equal(srcParse.Separators, dstParse.Separators) {
		return false
	}
	if !slices.Equal(srcParse.ExpenseKeywords, dstParse.ExpenseKeywords) {
		return false
	}
	if !slices.Equal(srcParse.IncomeKeywords, dstParse.IncomeKeywords) {
		return false
	}
	if !slices.Equal(srcParse.TodayKeywords, dstParse.TodayKeywords) {
		return false
	}
	if !slices.Equal(srcParse.YesterdayKeywords, dstParse.YesterdayKeywords) {
		return false
	}
	return true

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
	MaxGetCount  int64 `yaml:"maxGetCount"`
	ParserConfig `yaml:"parserConfig"`
}

type ParserConfig struct {
	Separators        []string `yaml:"separators"`
	ExpenseKeywords   []string `yaml:"expenseKeywords"`
	IncomeKeywords    []string `yaml:"incomeKeywords"`
	TodayKeywords     []string `yaml:"todayKeywords"`
	YesterdayKeywords []string `yaml:"yesterdayKeywords"`
}

func InitBizConfig() error {
	c := initBizConfig()
	err := util.EmptyCheck(c)
	if err != nil {
		return err
	}
	if !GetBizConfig().Equal(c) {
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

func SetBizConfig(c Config) {
	conf.Store(&c)
}
