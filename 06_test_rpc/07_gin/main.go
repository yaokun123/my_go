package main

import (
	pb "07_gin/proto"
	context2 "context"
	"github.com/gin-gonic/gin"
	"github.com/asim/go-micro/v3"
	"log"
)

func main() {
	// 初始化引擎
	gin.SetMode(gin.DebugMode)

	// 1、初始化路由
	router := gin.Default()
	// 注册异常捕获组件
	router.Use(gin.Recovery())

	// 2、做路由匹配
	router.GET("/", handle)

	// 3、运行
	router.Run(":9998")

}

func handle(context *gin.Context) {
	// context.Writer.WriteString("hello world")

	// 与micro对接(使用默认的mdns)

	// 1、初始化客户端
	srv := micro.NewService(
		// micro.Name("go.micro.web.bj38"),	// ???
		)
	srv.Init()

	microClient := pb.NewBj38Service("bj38", srv.Client())

	// 2、调用远程服务
	resp, err := microClient.Call(context2.TODO(), &pb.Request{
		Name: "demo",
	})
	if err != nil {
		log.Fatal("call rpc error!!!")
	}

	context.Writer.WriteString(resp.Msg)
}
