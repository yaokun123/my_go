package main

import (
	"os"
	"fmt"
	"net"
	"time"
)

func main()  {
	if len(os.Args) <=1{
		fmt.Println("usage:go run client3.go YOUR_CONTENT")
		return
	}

	fmt.Println("begin dial")
	conn,err := net.Dial("tcp","127.0.0.1:8686")
	if err != nil{
		fmt.Println("dial error:",err)
		return
	}

	defer conn.Close()
	fmt.Println("dial ok")

	time.Sleep(2*time.Second)
	data := os.Args[1]
	conn.Write([]byte(data))

	time.Sleep(time.Second * 10000)
}
