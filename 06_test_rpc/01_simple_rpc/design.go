package main

import (
	"log"
	"net/rpc"
)

const (
	SERVERNAME = "hello"
)

// 要求，服务端在注册rpc对象时，能让编译器检测出注册对象时合法的

// 创建接口，在接口中定义方法原型

type MyInterface interface {
	HelloWorld(name string, resp *string) error
	HelloWorld2(name string, resp *string) error
	HelloWorld3(name string, resp *string) error
}

// 包装注册接口，要求注册的参数必须实现MyInterface接口
func RegisterService(i MyInterface) {
	err := rpc.RegisterName(SERVERNAME, i)
	if err != nil {
		log.Fatal("注册rpc服务失败！")
	}
}

// =============================================

// 包装客户端
// 因为Call绑定在rpc.CLient对象上
type MyClient struct {
	c *rpc.Client
}

// 实现函数，原型参照上面的interface实现
func (this *MyClient) HelloWorld(name string, resp *string) error {
	serviceMethod := SERVERNAME + ".HelloWorld"
	return this.c.Call(serviceMethod, name, resp)
}
func (this *MyClient) HelloWorld2(name string, resp *string) error {
	serviceMethod := SERVERNAME + ".HelloWorld2"
	return this.c.Call(serviceMethod, name, resp)
}
func (this *MyClient) HelloWorld3(name string, resp *string) error {
	serviceMethod := SERVERNAME + ".HelloWorld3"
	return this.c.Call(serviceMethod, name, resp)
}
