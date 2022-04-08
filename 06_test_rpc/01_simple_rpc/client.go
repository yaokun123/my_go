package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 1、用rpc连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		log.Fatal("rpc.Dial error！")
	}
	defer conn.Close()

	// 2、调用远程函数
	name := "demo"
	var resp string  // 接受返回值	-- 传出参数
	var resp2 string // 接受返回值	-- 传出参数

	// 客户端在调用的时候容易将服务名.方法名写错，也可以包装一下客户端的调用
	//err = conn.Call("hello.HelloWorld", name, &resp)
	//err = conn.Call("hello.HelloWorld2", name, &resp2)

	// 客户端可以是使用包装的函数
	client := MyClient{
		c: conn,
	}
	err = client.HelloWorld(name, &resp)
	err = client.HelloWorld2(name, &resp2)
	if err != nil {
		log.Fatal("conn.Call error！")
	}
	fmt.Println(resp)
	fmt.Println(resp2)
}
