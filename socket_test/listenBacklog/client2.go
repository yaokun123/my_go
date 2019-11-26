package main

import (
	"net"
	"log"
	"time"
)

func main()  {
	var sl []net.Conn
	for i:=1;i<1000;i++{
		conn := establishConn(i)
		if conn != nil{
			sl = append(sl, conn)
		}
	}

	time.Sleep(10000*time.Second)
}

/**
对方服务的listen backlog满
还有一种场景就是对方服务器很忙，瞬间有大量client端连接尝试向server建立，server端的listen backlog队列满，
server accept不及时（即便不accept，那么在backlog数量范畴里面，connect都是会成功的，因为new conn已经加入到
server side的listen queue中了，accept只是从queue中取出一个conn而已，这将导致client端Dial阻塞）
 */
func establishConn(i int) net.Conn {
	conn,err := net.Dial("tcp","127.0.0.1:8686")
	if err != nil{
		log.Printf("%d:dial error:%s",i,err)
		return nil
	}
	log.Println(i,":connect to server ok")
	//	2019/11/26 10:02:09 1 :connect to server ok
	//	2019/11/26 10:02:09 2 :connect to server ok
	//	2019/11/26 10:02:09 3 :connect to server ok
	//	......
	//	2019/11/26 10:02:09 126 :connect to server ok
	//	2019/11/26 10:02:09 127 :connect to server ok
	//	2019/11/26 10:02:09 128 :connect to server ok
	//	......
	//	2019/11/26 10:02:09 129 :connect to server ok
	//	2019/11/26 10:02:22 130 :connect to server ok
	//	2019/11/26 10:02:35 131 :connect to server ok
	//	......
	return conn
}

//可以看出Client初始时成功地一次性建立了128个连接，然后后续每阻塞近10s才能成功建立一条连接。也就是说在server端
//backlog满时（未及时accept），客户端将阻塞在Dial上，直到server端进行一次accept。至于为什么是128，这与darwin
//下默认设置有关
//$sysctl -a|grep kern.ipc.somaxconn
//kern.ipc.somaxconn:128



//如果我们在ubuntu14.04上运行上述程序，我们的client端初始可以成功建立499条连接。
//如果server一直不accept，client端会一直阻塞么？我们去掉accept后的结果是：在darwin下，client端大约1分钟才会返回
//timeout
//	2019/11/26 10:29:42 129:dial error:dial tcp 127.0.0.1:8686: connect: operation timed out
//而如果server运行在ubuntu14.04上，client似乎一直阻塞我等了10分钟依旧没有返回。阻塞与否看来与server端的网络实现
//和设置有关