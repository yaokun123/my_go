package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main()  {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	defer client.Close()

	pong,_ := client.Ping(context.Background()).Result()
	fmt.Println(pong)


	err := client.Set(context.Background(),"demo","demo",0).Err()
	fmt.Println(err)
	res,_ := client.Get(context.Background(),"demo").Result()
	fmt.Println(res)
}