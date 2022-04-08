package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// 定义类对象
type World struct {
}

// 绑定类方法 - 1
// 如果不按照既定的规则来定义方法，也是可以编译通过的，只是在运行期间会报错
// 如何将报错提前到编译期？这就需要再进行封装-使用接口。MyInterface
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = "hello " + name
	return nil
}

// 绑定类方法 - 2
func (this *World) HelloWorld2(name string, resp *string) error {
	*resp = "hello2 " + name
	return nil
}

// 绑定类方法 - 3
// 这里没有按照既定的规则来定义类方法，但是可以通过编译，因为在语法上没有任何问题
// 会在程序运行期间报错
func (this *World) HelloWorld3(name string, resp *string) error {
	*resp = "hello3 "
	return nil
}

/**
RPC服务端
RPC使用了go语言特有的数据序列化gob，其他编程语言不能解析
可使用通用的序列化、反序列化json/protobuf
net/jsonrpc:json版本的rpc
*/
func main() {
	// 1、注册RPC服务，绑定对象方法
	/*err := rpc.RegisterName("hello", new(World))
	if err != nil {
		log.Fatal("注册rpc服务失败！")
	}*/
	// 使用包装好的注册方法，限制注册的RPC对象，此时如果注册的RPC对象没有按照规则定义，会在编译期间报错
	RegisterService(new(World))

	// 2、设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Fatal("net.Listen error！")
	}
	defer listener.Close()
	fmt.Println("开始监听。。。")

	// 3、建立连接
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("listener.Accept error！")
	}
	defer conn.Close()
	fmt.Println("连接成功。。。")

	// 4、绑定服务
	rpc.ServeConn(conn)
}
