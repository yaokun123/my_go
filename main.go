package main

import (
	"fmt"
	"time"
)

func main() {
	currentTimeStamp := time.Now().Unix()

	fmt.Println("==========")
	fmt.Println(currentTimeStamp)
}
