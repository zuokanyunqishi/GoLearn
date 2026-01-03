package coding

import (
	"context"
	"fmt"
	"testing"
)

// 顺序依次打印 dog  cat fish 各100次 使用协程
func TestPrintWords(t *testing.T) {

	exitChan := make(chan struct{})

	dogChan := noticeChan(1)
	catChan := noticeChan(1)
	fishChan := noticeChan(1)

	ctx, cannel := context.WithCancel(context.Background())

	go func() {
		for i := 0; i < 3; i++ {
			<-exitChan
		}
		cannel()
	}()

	// 1->2->3mm

	go printWord("dog", catChan, dogChan, exitChan)
	go printWord("cat", fishChan, catChan, exitChan)
	go printWord("fish", dogChan, fishChan, exitChan)

	<-ctx.Done()

}

func noticeChan(len int) chan struct{} {
	return make(chan struct{}, len)
}

func printWord(word string, tellchan, revchan, exitchan chan struct{}) {

	var counter int
	if word == "dog" {
		// 开始生产 单词
		fmt.Println(word)
		counter++
		tellchan <- struct{}{}
	}

	for {
		if counter >= 100 {
			exitchan <- struct{}{}
			return
		}
		<-revchan
		fmt.Println(word)
		counter++
		tellchan <- struct{}{}

	}
}

// 并发模型
func TestMakeSushu(t *testing.T) {

	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

func Generate(ch chan int) {

	for i := 2; ; i++ {
		ch <- i
	}

}

func Filter(in chan int, out chan int, prime int) {
	for {
		i := <-in

		if i%prime != 0 {
			out <- i
		}
	}
}
