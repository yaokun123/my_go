package model

import (
	_ "github.com/go-sql-driver/mysql"
	. "gopingan/common"
	"time"
	"xorm.io/xorm"
)

var UcsEngine *xorm.Engine

func SetEngine() *xorm.Engine {
	server := Cfg.MustValue("db", "server", "127.0.0.1")
	username := Cfg.MustValue("db", "username", "root")
	password := Cfg.MustValue("db", "password", "pass")
	dbName := Cfg.MustValue("db", "db_name", "go_display")
	orm, err := xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+dbName+"?charset=utf8")
	PanicIf(err)
	//orm.TZLocation = time.Local
	orm.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	orm.ShowSQL(Cfg.MustBool("db", "show_sql", false))
	//orm.Logger = xorm.NewSimpleLogger(Log.GetWriter())
	return orm
}
