package main

import (
	"04_grpc/pb/person"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 1、连接rpc服务
	// WithInsecure	表示以安全的方式操作
	// grpcConn,err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	// grpcConn,err := grpc.Dial("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, err := grpc.Dial("127.0.0.1:8800")
	if err != nil {
		log.Fatal("grpc.Dial error")
	}
	defer grpcConn.Close()

	// 2、初始化grpc客户端
	grpcClient := person.NewSayNameClient(grpcConn)

	// 创建
	// 3、调用远程函数
	p := person.Teacher{
		Name: "demo",
		Age:  18,
	}
	r, err := grpcClient.SayHello(context.TODO(), &p)
	fmt.Println(r)

}
