package main

import (
	//"github.com/micro/go-micro/client"
	pb "07_gin/proto"
	context2 "context"
	"github.com/gin-gonic/gin"
	"github.com/micro/micro/v3/service"
	"log"
)

func main() {
	// 1、初始化路由
	router := gin.Default()

	// 2、做路由匹配
	router.GET("/", func(context *gin.Context) {
		// context.Writer.WriteString("hello world")

		// 与micro对接(使用默认的mdns)

		// 1、初始化客户端
		srv := service.New(
		//service.Name("bj38"),
		)
		microClient := pb.NewBj38Service("bj38", srv.Client())

		// 2、调用远程服务
		resp, err := microClient.Call(context2.TODO(), &pb.Request{
			Name: "demo",
		})
		if err != nil {
			log.Fatal("call rpc error!!!")
		}

		context.Writer.WriteString(resp.Msg)
	})

	// 3、运行
	router.Run(":9999")

}
