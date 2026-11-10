package redis

import (
	"context"
	"fmt"
	"go-production-stack/internal/config"

	goredis "github.com/redis/go-redis/v9"
)

var (
	Client *goredis.Client
	Ctx    = context.Background()
)

func InitRedis() {
	cfg := config.Load()

	address := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	Client = goredis.NewClient(&goredis.Options{
		Addr:     address,
		Password: cfg.RedisPassword,
	})
	fmt.Println("Redis Address", address)
}

func CloseRedis() {
	if Client != nil {
		if err := Client.Close(); err != nil {
			fmt.Println("Error Closing Redis: ", err)
			return
		}
		fmt.Println("Redis Connection Closed")
	}

}

// TODO: Add context timeout handling for Redis-dependent operations
// to prevent requests from hanging when Redis is unavailable.
