package main

import (
	"fmt"
	"sync"

	"github.com/go-basic/uuid"
)

var (
	lock  = sync.Mutex{}
	store = make(map[int]string)
)

// 协程 非抢占式 多任务处理,协程主动交出控制权
// race 数据冲突调试
func main() {
	//routine1()

	// 同步阻塞调用
	wait := sync.WaitGroup{}
	wait.Add(500)
	for i := 0; i < 500; i++ {
		go syncMutex(i, &wait)
	}

	// 等待全部协程运行完毕
	wait.Wait()

	for key, value := range store {
		fmt.Printf("store[%d]=%s\n", key, value)
	}

	//go func() {
	//
	//}()

	//time.Sleep(time.Minute)
}

func routine1() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println(`helloWord`, i)
			}
		}(i)
	}
}

func syncMutex(key int, wait *sync.WaitGroup) {

	lock.Lock()
	for i := 0; i < 500; i++ {
		store[key+i+1] = makeRandStr()

	}
	lock.Unlock()
	wait.Done()

}

func makeRandStr() string {
	return uuid.New()

}
