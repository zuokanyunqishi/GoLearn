package util

import (
	"github.com/go-redis/redis/v8"
	"log"
)

var Redis *redis.Client

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Username: "",
		Password: "",
		PoolSize: 3,
	})
	_, err := Redis.Ping(Redis.Context()).Result()
	if err == redis.Nil {
		log.Fatal("Redis异常", err)
	} else if err != nil {
		log.Fatal("失败:", err.Error())
	}

}
