package redisApplication

import (
	"context"
	"fmt"
	"github.com/go-basic/uuid"
	"github.com/syyongx/php2go"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestMakeBigKey(t *testing.T) {

	ctx := context.Background()
	redisC := RedisCluster()
	var wg sync.WaitGroup
	// 生产几百万的 key
	for i := 1800; i <= 2400; i++ {
		wg.Add(1)
		go func(c int) {

			size := 5000
			strings := make([]interface{}, size*2)

			start := 0
			for j := (c - 1) * size; j < size*c; j++ {
				strings[start] = j
				start = start + 1
				strings[start] = uuid.New()
				start = start + 1
			}
			fmt.Println(strings)
			redisC.HMSet(ctx, "testRedis:111", strings...)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func TestGetBigKey(t *testing.T) {

	for i := 0; i < 10000; i++ {
		go func() {
			for {
				get := RedisCluster().HGet(context.TODO(), "testRedis:111", strconv.Itoa(php2go.Rand(1, 500000)))
				fmt.Println(get.Result())
			}

		}()
	}

	time.Sleep(time.Second * 10)
}
