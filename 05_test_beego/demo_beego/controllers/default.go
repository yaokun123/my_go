package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"        // {{.Website}}
	c.Data["Email"] = "astaxie@gmail.com" // {{.Email}}

	// 默认支持 tpl 和 html 的后缀名，如果想设置其他后缀你可以调用 beego.AddTemplateExt 接口设置
	// beego 采用了 Go 语言默认的模板引擎，所以和 Go 的模板语法一样，Go 模板的详细使用方法请参考《Go Web 编程》
	c.TplName = "index.tpl"

	// beego 默认注册了 static 目录为静态处理的目录
	// 注册样式：URL 前缀和映射的目录（在/main.go文件中beego.Run()之前加入）：
	// StaticDir["/static"] = "static"

	// 用户可以设置多个静态文件处理目录，例如你有多个文件下载目录 download1、download2，
	// 你可以这样映射（在 /main.go 文件中 beego.Run() 之前加入）：
	// beego.SetStaticPath("/down1", "download1")
	// beego.SetStaticPath("/down2", "download2")
	// 这样用户访问 URL http://localhost:8080/down1/123.txt 则会请求 download1 目录下的 123.txt 文件
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	// appname := beego.AppConfig.String("appname")
	appname := beego.AppConfig.String("AppName")
	dbHost := beego.AppConfig.String("db::host")

	dbUser, _ := beego.GetConfig("String", "db::user", "dbuser")
	dbUserStr, ok := dbUser.(string)
	c.Ctx.WriteString("welcome user center!!" + " " + appname + " " + dbHost + " ")
	if ok {
		c.Ctx.WriteString(dbUserStr)
	}

	// 返回json
	// c.ServeJSON()

	// 跳转
	// c.Redirect("/", 302)
}
