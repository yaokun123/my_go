package main

import (
	"bj38/handler"
	pb "bj38/proto"
	//"github.com/asim/go-micro/v3"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	//"github.com/asim/go-micro/plugins/registry/consul/v3"
	//"github.com/asim/go-micro/v3/registry"
)

func main() {
	// 更改服务发现为consul
	/*reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:8500"}
	})*/

	// Create service
	srv := service.New(
		service.Name("bj38"),
	)

	// Register handler
	pb.RegisterBj38Handler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
