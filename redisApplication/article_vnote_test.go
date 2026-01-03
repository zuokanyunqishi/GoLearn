package redisApplication

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/syyongx/php2go"

	"github.com/golang-module/carbon/v2"

	"github.com/go-redis/redis/v8"
)

func TestInitData(t *testing.T) {

	ctx := context.Background()

	userId := 10600

	for i := 10000; i < 10500; i++ {
		//Redis().Del(ctx, fmt.Sprintf("%d", i))
		//continue
		now := carbon.Now()
		articleSign := fmt.Sprintf("article:%d", i)
		Redis().ZAdd(ctx, "time:article", &redis.Z{
			Score:  float64(now.Timestamp()),
			Member: articleSign,
		})

		Redis().ZAdd(ctx, "score:note", &redis.Z{
			Score:  0,
			Member: articleSign,
		})

		userSign := fmt.Sprintf("user:%d", userId)
		Redis().HSet(ctx, userSign, []string{
			"name", php2go.Uniqid(fmt.Sprintf("%d_", userId)),
		})

		Redis().HSet(ctx, articleSign, []string{
			"title", "xxxx",
			"body", "xxxxxxx",
			"author", "user:",
			"time", fmt.Sprintf("%s", now.String()),
		})

		userId += 1
		time.Sleep(time.Second)
	}

}

func TestVnote(t *testing.T) {
	ctx := context.Background()

	Redis().ZIncrBy(ctx, "score:note", 3, "article:10003")
}
