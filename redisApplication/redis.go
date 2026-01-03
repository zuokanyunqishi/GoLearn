package redisApplication

import (
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var redisCluster *redis.ClusterClient

func Redis() *redis.Client {
	return redisClient
}

func RedisCluster() *redis.ClusterClient {
	return redisCluster
}
func init() {
	//
	//redisClient = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Username: "",
	//	Password: "",
	//	PoolSize: 3,
	//})
	//_, err := redisClient.Ping(redisClient.Context()).Result()
	//if err == redis.Nil {
	//	log.Fatal("Redis异常", err)
	//} else if err != nil {
	//	log.Fatal("失败:", err.Error())
	//}

	redisCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.0.103:6381", "192.168.0.103:6382", "192.168.0.103:6383", "192.168.0.103:6384", "192.168.0.103:6385", "192.168.0.103:6386"},
	})
}
