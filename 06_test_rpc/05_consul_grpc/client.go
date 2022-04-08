package main

import (
	"05_consul_grpc/pb/person"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	// 1、初始化consul配置
	consulConfig := api.DefaultConfig()

	// 2、创建consul对象（可以指定consul属性 IP/Port）
	consulClient, err := api.NewClient(consulConfig)

	// 3、服务发现，从consul上获取健康的服务
	// service	服务名  注册服务时指定的
	// tag		别名		任选一个
	// passingOnly	是否通过健康检查，true
	// q	查询参数 通常穿nil
	// 返回值
	// []*ServiceEntry 存储服务的切片，最好判断是否为空
	// *QueryMeta	额外查询返回的值，一般为nil
	// error
	ServiceEntry, _, err := consulClient.Health().Service("grpc And Consul", "consul", true, nil)
	if err != nil {
		log.Fatal("获取微服务失败")
	}

	if len(ServiceEntry) == 0 {
		log.Fatal("无可用的地址")
	}

	// 组装要请求的微服务地址
	addr := ServiceEntry[0].Service.Address + ":" + strconv.Itoa(ServiceEntry[0].Service.Port)

	//////////////////////////////以下为grpc服务远程调用///////////////////////////////////////////////

	// 1、连接rpc服务
	// WithInsecure	表示以安全的方式操作
	// grpcConn,err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc.Dial error")
	}
	defer grpcConn.Close()

	// 2、初始化grpc客户端
	grpcClient := person.NewHelloClient(grpcConn)

	// 创建
	// 3、调用远程函数
	p := person.Person{
		Name: "demo",
		Age:  18,
	}
	r, err := grpcClient.SayHello(context.TODO(), &p)
	fmt.Println(r.Name)
}
