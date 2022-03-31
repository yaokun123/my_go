package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	currentTime := time.Now()
	fmt.Printf("%T,%#v,%v", currentTime, currentTime, currentTime)
	fmt.Println("========================")
	fmt.Println(currentTime.Year(), currentTime.YearDay())
	fmt.Println(currentTime.Month())
	fmt.Println(currentTime.Day())
	fmt.Println(currentTime.Hour())
	fmt.Println(currentTime.Minute())
	fmt.Println(currentTime.Second())

	// unix时间戳
	fmt.Println(currentTime.Unix())
	fmt.Println(currentTime.UnixNano())

	// 字符串格式
	/**
	占位		说明
	2006 	4位数字的年
	01		2位数字的月
	02		2位数字的日
	03		12进制的时
	15		24进制的时
	04		2位数字的分钟
	05		2位数字的秒
	*/
	fmt.Println(currentTime.Format("2006-01-02 15:04:05"))
}
