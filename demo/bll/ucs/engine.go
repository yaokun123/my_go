package ucs

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
)

var UcsEngine *xorm.Engine

func init() {
	e, err := xorm.NewEngine("mysql", "root:phpdev@192.168.109.245/ucs?charset=utf8")

	if err != nil {
		log.Fatal("引擎初始化失败")
	}

	err = e.Ping()
	if err != nil {
		log.Println("连接创建失败")
	}
	e.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	e.SetMaxIdleConns(2)
	e.SetMaxOpenConns(10)

	UcsEngine = e
}
