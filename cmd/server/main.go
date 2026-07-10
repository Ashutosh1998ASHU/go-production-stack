package main

import (
	"fmt"
	"go-production-stack/internal/config"
	"go-production-stack/internal/handlers"
	"go-production-stack/internal/redis"
	"net/http"
)

// var ctx = context.Background()

// var rdb = redis.NewClient(&redis.Options{
// 	Addr: "redis:6379",
// })

func main() {
	cfg := config.Load()
	redis.InitRedis()
	http.HandleFunc("/", handlers.HomeHandler)
	fmt.Printf("Starting server on :%s\n", cfg.ServerPort)
	err := http.ListenAndServe(":"+cfg.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}
