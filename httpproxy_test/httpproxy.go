package main

import (
	"net"
	"log"
	"fmt"
	"bytes"
	"net/url"
	"strings"
	"io"
)

func main()  {
	log.SetFlags(log.LstdFlags|log.Lshortfile)

	//启动代理监听
	l,err := net.Listen("tcp",":8889")
	if err != nil{
		log.Panic(err)
	}

	//监听接收代理请求
	for{
		client,err := l.Accept()
		if err != nil{
			log.Panic(err)
		}
		
		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn)  {
	var b [1024]byte
	n,err := client.Read(b[:])
	if err != nil{
		log.Panic(err)
		return
	}
	var method,host,address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:],'\n')]),"%s%s",&method,&host)
	hostPortURL,err := url.Parse(host)
	if err != nil{
		log.Panic(err)
		return
	}

	if hostPortURL.Opaque == "443"{//https访问
		address = hostPortURL.Scheme+":443"
	}else{//http访问
		if strings.Index(hostPortURL.Host,":") == -1{//host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		}else{
			address = hostPortURL.Host
		}
	}

	//代理服务器和远程服务器建立连接
	//获得了请求的host和port，就开始拨号吧
	server,err := net.Dial("tcp",address)
	if err != nil{
		log.Panic(err)
		return
	}

	//数据转发
	//拨号成功后，就可以进行数据代理传输了
	if method == "CONNECT"{
		fmt.Fprint(client,"HTTP/1.1 200 connection established\r\n\r\n")
	}else{
		server.Write(b[:n])
	}

	//进行转发
	go io.Copy(server,client)
	io.Copy(client,server)
}
