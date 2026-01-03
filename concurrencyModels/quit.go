package concurrencyModels

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(str string, quit chan string) (c chan string) {
	c = make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s: %d", str, i):
			// do something
			case <-quit:
				// 退出前的清理动作
				// clean up()
				quit <- "well done"
				return
			}
		}
	}()

	return

}

// RunQuit1 退出模式
func RunQuit1() {
	quit := make(chan string)
	c := boring("xiaoming", quit)

	rand.Seed(time.Now().Unix())
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "bye"
	fmt.Println(<-quit)

}
