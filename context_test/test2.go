package main

import (
	"time"
	"fmt"
)

//注意因为都是main包，test1,test2方法在test.go中已经存在。所以这里使用test3方法
func test3()  {
	go func() {
		for range time.Tick(time.Second){
			fmt.Println("孙子goroutine中打印的东西");
		}
	}()
	fmt.Println("子goroutine结束")
}

func main()  {
	go test3()

	time.Sleep(10*time.Second)

	fmt.Println("main stop")
}

//我们在test3 子goroutine中又启动了一个孙子goroutine
//按照我们的想象，test3子goroutine结束，孙子goroutine也会跟着结束，但是结果并没有，孙子goroutine继续运行
//这是因为，孙子goroutine也是依赖于main goroutine的

//那么我们如果想让子goroutine结束孙子goroutine也我跟结束的话
//就可以使用context上下文
//在子goroutine中定义一个上下文，孙子goroutine中判断上下文状态（见test3.go）