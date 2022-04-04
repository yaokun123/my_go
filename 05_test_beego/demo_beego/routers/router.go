package routers

import (
	"context"
	"demo_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 基础路由
	beego.Get("/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Post("/post", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	beego.Any("/any", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})

	// 正则路由

	// 自动路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
