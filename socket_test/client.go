package main

import (
	"net"
	"fmt"
)

func main()  {
	conn,err := net.Dial("tcp","127.0.0.1:8866")//阻塞Dial
	//conn1,err1 := net.Dial("tcp","127.0.0.1:8866",2*time.Second)//带上超时机制的Dial

	if err != nil{
		//handle error
		fmt.Println("client dial error:",err)
	}

	defer  conn.Close()
	//read or write on conn

}
