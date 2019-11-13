package main

import (
	"time"
	"fmt"
	"context"
)

func test4(){
	//在子goroutine中生成一个context上下文
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for range time.Tick(time.Second){
			select {
			case <-ctx.Done():
				fmt.Println("子goroutine结束，孙子goroutine通过判断上下文状态也要结束")
				return
			default:
				fmt.Println("孙子goroutine中打印的东西")
			}

		}
	}()
	fmt.Println("子goroutine结束")
}
func main()  {
	go test4()

	time.Sleep(10*time.Second)

	fmt.Println("main stop")
}

//下面介绍一下context的使用：
