package main

import (
	"os"
	"fmt"
)

func main()  {
	/**
	1、将字符串写入文件
	 */
	 file := "/tmp/log.txt"
	 f,err := os.OpenFile(file,os.O_RDWR|os.O_APPEND|os.O_CREATE,0666)
	 if err != nil{
	 	fmt.Println(err)
	 	return
	 }

	 /*l,e := f.WriteString("Hello World\r\n")
	 if e != nil{
	 	fmt.Println(e)
	 }
	 fmt.Println(l)*/



	 /**
	 2、将字节写入文件
	  */
	/*d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
	n2, err := f.Write(d2)
	fmt.Println(n2)*/


	/**
	3、将字符串一行一行的写入文件
	 */
	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

	for _,v := range d{
		fmt.Fprintln(f,v)
	}
}

