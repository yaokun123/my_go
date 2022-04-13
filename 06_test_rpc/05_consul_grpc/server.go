package main

import (
	"05_consul_grpc/pb/person"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 定义类
type Children struct {
}

// 按接口绑定类方法
func (this *Children) SayHello(ctx context.Context, t *person.Person) (*person.Person, error) {
	t.Name += " is Sleeping"
	return t, nil
}

func main() {
	// 把grpc注册到

	// 1、初始化consul配置
	consulConfig := api.DefaultConfig()

	// 2、创建consul对象
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Fatal("api.NewClient error")
	}

	// 3、告诉consul，即将注册的服务配置
	reg := api.AgentServiceRegistration{
		ID:      "bj38",
		Tags:    []string{"grpc", "consul"},
		Name:    "grpc And Consul",
		Address: "127.0.0.1",
		Port:    8800,
		Check: &api.AgentServiceCheck{
			CheckID:  "consul grpc test",
			TCP:      "127.0.0.1:8800",
			Timeout:  "1s",
			Interval: "5s",
		},
	}

	// 4、注册grpc服务到consul上
	err = consulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal("consul error")
	}

	// 注销服务
	// consulClient.Agent().ServiceDeregister("bj38")

	//////////////////////////////以下为grpc服务远程调用///////////////////////////////////////////////

	// 1、初始一个grpc对象
	grpcServer := grpc.NewServer()

	// 2、注册服务
	person.RegisterHelloServer(grpcServer, new(Children))

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
