package routers

import (
	"demo_beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 1、基础路由
	/*beego.Get("/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Post("/post", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	beego.Any("/any", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})*/

	// 正则路由

	// 2、固定路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
