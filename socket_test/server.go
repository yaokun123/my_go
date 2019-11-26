package main

import (
	"net"
	"fmt"
)

func main()  {
	l,err := net.Listen("tcp","127.0.0.1:8866")
	if err != nil{
		fmt.Println("server listen error:",err)
		return
	}
	defer l.Close()

	for{
		c,err := l.Accept()
		if err != nil{
			fmt.Println("server accept error:",err)
			break
		}

		go handleConn(c)
	}
}

func handleConn(c net.Conn)  {
	defer c.Close()
	for{
		//read from the connection
		//... ...

		fmt.Println("=============")

		//write to the connrction
		//... ...
	}
}
