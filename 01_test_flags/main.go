package main

import (
	"flag"
	_ "flag"
	"fmt"
	"os"
)

/** cobar第三方库*/

func main() {
	// 一般简单的参数传递使用os.Args即可
	args := os.Args
	fmt.Println(args)

	// 复杂一点的参数，例如-H -P 等需要使用flag包
	var (
		host string
		port int
		h    bool
		help bool
	)
	// 	解析变量的指针，命令行中指定的参数名，默认值，帮助
	flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	host2 := flag.String("H", "127.0.0.1", "连接地址") // 另一种方法，返回指针类型

	flag.IntVar(&port, "P", 22, "连接端口")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&help, "help", false, "帮助")
	fmt.Println(host2)

	// 重写帮助函数
	/*flag.Usage = func() {
		fmt.Println("uasge: testflag [-H 127.0.0.1] [-P 22]")
		flag.PrintDefaults()
	}*/

	// 解析参数
	flag.Parse()

	if h || help {
		flag.Usage()
		return
	}

	fmt.Println(host, port, h, help)
	fmt.Println(flag.NArg()) // 用于接受其他没有指定名称的参数
	fmt.Println(flag.Args()) // 用于接受其他没有指定名称的参数
}
