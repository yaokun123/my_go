package main

import (
	"log"
	"os"
)

func main() {

	//使用非常简单，函数名字和用法也和fmt包很相似，但是它的输出默认带了时间戳。
	//这样我们很清晰的就知道了，记录这些日志的时间，这对我们排查问题，非常有用。
	log.Println("这是一个测试", "hello world")
	log.Printf("这是一个测试%s", "hello world")

	//有了时间了，我们还想要更多的信息，比如发生的源代码行号等
	//对此日志包log 为我们提供了可定制化的配制，让我们可以自己定制日志的抬头信息
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("这是一个加上源代码行号的测试")
	/*const (
		Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
		Ltime                         // the time in the local time zone: 01:23:23
		Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
		Llongfile                     // full file name and line number: /a/b/c/d.go:23
		Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
		LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
		LstdFlags     = Ldate | Ltime // initial values for the standard logger
	)*/

	//我们大部分情况下，都有很多业务，每个业务都需要记录日志，那么有没有办法，能区分这些业务呢？
	// 这样我们在查找日志的时候，就方便多了。
	//对于这种情况，Go语言也帮我们考虑到了，这就是设置日志的前缀，比如一个用户中心系统的日志，我们可以这么设置。
	log.SetPrefix("[userCenter]")
	log.Println("这是一个有前缀的测试")

	//log包除了有Print系列的函数，还有Fatal以及Panic系列的函数
	//其中Fatal表示程序遇到了致命的错误，需要退出，这时候使用Fatal记录日志后，然后程序退出，
	// 也就是说Fatal相当于先调用Print打印日志，然后再调用os.Exit(1)退出程序。

	//同理Panic系列的函数也一样，表示先使用Print记录日志，然后调用panic()函数抛出一个恐慌，
	// 这时候除非使用recover()函数，否则程序就会打印错误堆栈信息，然后程序终止。

	// 设置输出到文件（默认输出到终端）
	file := "/tmp/log.txt"
	of, _ := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	log.SetOutput(of)
	log.Println("测试打印到文件中")

	//实现原理:
	/**
	func New(out io.Writer, prefix string, flag int) *Logger {
		return &Logger{out: out, prefix: prefix, flag: flag}
	}
	var std = New(os.Stderr, "", LstdFlags)

	从以上源代码可以看出，变量std其实是一个*Logger，通过log.New函数创建，
	默认输出到os.Stderr设备，前缀为空，日志抬头信息为标准抬头LstdFlags。

	os.Stderr对应的是UNIX里的标准错误警告信息的输出设备，同时被作为默认的日志输出目的地。
	除此之外，还有标准输出设备os.Stdout以及标准输入设备os.Stdin。
	*/
}
