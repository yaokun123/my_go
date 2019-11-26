package main

import (
	"net"
	"log"
	"time"
)

func main()  {
	l,err := net.Listen("tcp","127.0.0.1:8686")
	if err != nil{
		log.Println("error listen:",err)
		return
	}

	defer l.Close()
	log.Println("listen ok")//	2019/11/26 10:01:54 listen ok

	var i int
	for{
		//每10秒钟accept一次
		time.Sleep(10*time.Second)
		if _,err := l.Accept();err != nil{
			log.Println("accept error:",err)
			break
		}
		i++
		log.Printf("%d:accept a new connection\n",i)
		// 	2019/11/26 10:02:09 1:accept a new connection
		//	2019/11/26 10:02:19 2:accept a new connection
		//	2019/11/26 10:02:29 3:accept a new connection
		//	2019/11/26 10:02:39 4:accept a new connection
	}
}
