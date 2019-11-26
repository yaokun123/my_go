package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

func main()  {
 /**
 1、将整个文件读取到内存
 将整个文件读取到内存是最基本的文件操作之一。这需要使用 ioutil 包中的 ReadFile 函数。
  */

  file := "/tmp/log.txt"
  data,err := ioutil.ReadFile(file)
  if err != nil{
  	fmt.Println("file reading error:",err)
  	return
  }
  fmt.Println("Contents of file",string(data))


  /**
  2、分块读取文件
  我们学习了如何把整个文件读取到内存。当文件非常大时，尤其在 RAM 存储量不足的情况下，把整个文件都读入内存是没有意义的。
  更好的方法是分块读取文件。这可以使用 bufio 包来完成。
   */
   f,_ := os.Open(file)
   defer f.Close()

   r := bufio.NewReader(f)
   b := make([]byte,3)
   var i int
   for{
   	_,err := r.Read(b)
   	if err != nil{
   		break
	}
   	fmt.Println(i,string(b))
   	i++
   }


   /**
   3、逐行读取文件
   这可以使用 bufio 来实现。
    */
	f2,e := os.Open(file)
	if e != nil{
		fmt.Println("read failed")
		return
	}
	defer f2.Close()
    s := bufio.NewScanner(f2)
    for s.Scan(){
    	fmt.Println(s.Text())
	}
}
