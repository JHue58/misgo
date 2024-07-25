package db

import (
	"github.com/jhue/misgo/db/model"
	"github.com/jhue/misgo/internal/conf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var ins *gorm.DB

func initDB() {
	d, err := newDB()
	if err != nil {
		panic(err)
	}
	ins = d
	err = autoMigrate(d)
	if err != nil {
		panic(err)
	}
}

func Get() *gorm.DB {
	if ins == nil {
		initDB()
	}
	return ins
}

func newDB() (*gorm.DB, error) {
	config := conf.GetConfig()

	db, err := gorm.Open(sqlite.Open(config.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(config.IdleConn)                                    // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(config.MaxConn)                                     // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Minute) // 设置连接的最大生命周期

	return db, nil
}

func autoMigrate(d *gorm.DB) (err error) {
	err = d.AutoMigrate(model.GetEmptyModels()...)
	if err != nil {
		return err
	}
	return
}
