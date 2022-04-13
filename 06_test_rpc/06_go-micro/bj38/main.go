package main

import (
	"bj38/handler"
	pb "bj38/proto"

	//"github.com/micro/micro/v3/service"
	//"github.com/micro/micro/v3/service/logger"
	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	// Create service
	// micro/micro 创建服务
	/*srv := service.New(
		service.Name("bj38"),
	)*/

	// asim/go-micro 创建服务
	srv := service.NewService(
		service.Name("bj38"),
		service.Version("latest"),
		)

	// Register handler
	pb.RegisterBj38Handler(srv.Server(), handler.New())
	//pb.RegisterBj38Handler(srv.Server(), new(handler.Bj38))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
