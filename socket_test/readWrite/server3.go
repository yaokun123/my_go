package main

import (
	"net"
	"fmt"
)

func main()  {
	l,err := net.Listen("tcp","127.0.0.1:8686")
	if err != nil{
		fmt.Println("listen error:",err)
	}
	defer l.Accept()

	for{
		conn,err := l.Accept()
		if err != nil{
			fmt.Println("accept error:",err)
		}

		go handleConn(conn)
	}
}
func handleConn(conn net.Conn)  {
	defer conn.Close()

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
socket中有部分数据：
如果socket中有部分数据，且长度小于一次read操作所期望读出的数据长度，那么Read将会成功读出这部分数据并返回，
而不是等待所有期望数据全部读取后再返回。
我们通过client3.go发送hi到server端运行结果：（go run client3.go hi）
start to read from conn
read 2 bytes,content is hi

Client向socket中写入了两个字节数据（hi），server端创建了一个len=10的slice，等待Read将读取的数据放入slice
server随后读取到那两个字节"hi"。Read成功返回，n=2,err=nil
 */


 /**
 socket中有足够数据
 如果socket中有数据，且长度大于等于一次Read操作所期望读出的数据长度，那么Read将会成功读出这部分数据并返回。这个
 情景是最符合我们对Read的期待：Read将用socket中的数据将我们传入的slice填满后返回:n=10,err=nil。

 我们通过client3.go向server3发送如下内容：abcdefghij12345：（go run client3.go abcdefghij12345）
 start to read from conn
 read 10 bytes,content is abcdefghij
 start to read from conn
 read 5 bytes,content is 12345
 start to read from conn

 client端发送的内容长度为15个字节，Server端Read buffer的长度为10，因此Server Read第一次返回时只会读取10个字节;
 socket中还剩余5个字节数据，Server再次Read时会把剩余数据读出
  */