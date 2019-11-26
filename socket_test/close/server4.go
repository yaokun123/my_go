package main

import (
	"net"
	"fmt"
	"time"
)

func main()  {
	l,err := net.Listen("tcp","127.0.0.1:8686")
	if err != nil{
		fmt.Println("listen error:",err)
	}

	defer l.Close()

	for{
		conn,err := l.Accept()
		if err != nil{
			fmt.Println("accept error:",err)
		}

		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn)  {
	defer conn.Close()
	time.Sleep(5*time.Second)

	for{
		//read from the connection
		var buf = make([]byte,10)
		fmt.Println("start to read from conn")
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn read error:",err)
			return
		}

		fmt.Printf("read %d bytes,content is %s\n",n,string(buf[:n]))
	}

}


/**
Socket关闭
如果client端主动关闭了socket，那么server的Read将会读到什么呢？这里分为"有数据关闭"和"无数据关闭"。

有数据关闭是指再client关闭时，socket中还有server端未读取的数据
我们通过client4.go向server4发送如下内容：abcdefghij12345：（go run client4.go abcdefghij12345）
start to read from conn
read 10 bytes,content is abcdefghij
start to read from conn
read 5 bytes,content is 12345
start to read from conn
conn read error: EOF

从输出结果来看，当client端close socket退出后，server4依旧没有开始Read，5s后第一次Read成功读取数据abcdefghij
第二次Read12345，第三次Read时，由于client端socket关闭，Read返回EOF error
通过上面这个例子，我们也可以猜测出"无数据关闭"情景下的结果，那就是Read直接返回EOF error
 */
