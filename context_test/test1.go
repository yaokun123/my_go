package main

import (
	"time"
	"fmt"
	"runtime"
)


//总结：由main goroutine产生的所有子goroutine都会伴随main的结束而结束
//如果main goroutine没结束，子goroutine有两种情况：
//情况一：子goroutine在main goroutine结束之前自己就已经结束了，此时只有main goroutine
//情况二：子goroutine在main goroutine结束之前也未结束，此时有main goroutine和子goroutine
func test1(){
	for range time.Tick(time.Second){
		fmt.Println("test1 goroutine")
	}
}

func test2()  {
	fmt.Println("test2 中共有协程：",runtime.NumGoroutine())//main + test1 + test2
	fmt.Println("test2 goroutine")
}

func main()  {
	go test1()
	go test2()
	fmt.Println("main 中共有协程：",runtime.NumGoroutine())//main + test1 + test2

	time.Sleep(10*time.Second)

	fmt.Println("main 中共有协程：",runtime.NumGoroutine())//main + test2
	fmt.Println("main stop")
}

//test2 goroutine只打印一次
//test1 无限循环打印（每秒打印一次）

//test1 goroutine 和 test2 goroutine都会随着main goroutine的结束而结束
//而如果main goroutine是永不结束的话,test2 goroutine会自动结束,test1 goroutine不会自动结束
