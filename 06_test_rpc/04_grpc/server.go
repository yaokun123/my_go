package main

import (
	person "04_grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 定义类
type Children struct {
}

// 按接口绑定类方法
func (this *Children) SayHello(ctx context.Context, t *person.Teacher) (*person.Teacher, error) {
	t.Name += "is Sleeping"
	return t, nil
}

func main() {
	// 1、初始一个grpc对象
	grpcServer := grpc.NewServer()

	// 2、注册服务
	person.RegisterSayNameServer(grpcServer, new(Children))

	// 3、设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Fatal("net.Listen error！")
	}
	defer listener.Close()
	fmt.Println("开始监听。。。")

	// 4、不需要在建立连接，这些都由底层做了。启动服务即可
	grpcServer.Serve(listener)
}
