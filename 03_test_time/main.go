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

	// 生成时间
	year, month, day := 2022, 3, 31
	time2022 := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	fmt.Printf("%T,%#v,%v\n", time2022, time2022, time2022)

	time2022 = time.Unix(1648693364, 0)
	fmt.Printf("%T,%#v,%v\n", time2022, time2022, time2022)

	time2022, err := time.Parse("2006-01-02 15:04:05", "2022-03-31 10:22:44")
	if err == nil {
		fmt.Printf("%T,%#v,%v\n", time2022, time2022, time2022)
	}

	time2022, err = time.ParseInLocation("2006-01-02 15:04:05", "2022-03-31 10:22:44", time.Local)
	if err == nil {
		fmt.Printf("%T,%#v,%v\n", time2022, time2022, time2022)
	}

	// 时间差
	duration := time.Since(time2022) // now - since
	fmt.Printf("%T,%#v,%v\n", duration, duration, duration)
	fmt.Println(duration.Seconds())

	duration = time.Until(time2022) // since - now
	fmt.Printf("%T,%#v,%v\n", duration, duration, duration)

	tomDuration, err := time.ParseDuration("24h") // 一天以后
	tom := time.Now().Add(tomDuration)
	fmt.Printf("%T,%#v,%v\n", tom, tom, tom)
}
