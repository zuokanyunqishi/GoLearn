package redisApplication

import "context"

func hyperLog() {
	ctx := context.Background()
	RedisCluster().PFAdd(ctx, "hyperLog", 1, 2, 1, 1, 11)
}
