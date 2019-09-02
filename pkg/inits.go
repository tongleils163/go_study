package pkg

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pmars/beego/logs"
)

var Engine *xorm.Engine
var EngineRecognize *xorm.Engine
var EngineTwo *xorm.Engine

func initMysqlEngine() {
	logs.Debug("Init Mysql Engine Now...")
	// 初始化Mysql
	Engine = initOneMysqlEngine(
		"cu:password@tcp(172.16.18.26:3306)/cread?timeout=3s&parseTime=true&loc=Local&charset=utf8mb4",
		500,
		20)
	logs.Debug("Init Mysql Engine Done!!!")

}

func initOneMysqlEngine(conn string, maxActive, maxIdle int) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(maxActive)
	engine.SetMaxIdleConns(maxIdle)
	if err = engine.Ping(); err != nil {
		logs.Info("mysql error:%v", err)
		panic(err)
	}

	return engine
}
