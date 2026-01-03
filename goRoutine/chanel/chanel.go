package main

import (
	"fmt"
	"time"
)

func main() {

	//var achan chan string
	// send or receive  nil chan deadlock
	//achan<-"s"
	//<-achan
	// close nil chan , panic
	// close(achan)

	//
	//cchan := make(chan int,10)
	//close(cchan)
	//// receive from close chan , it is zero value
	//fmt.Println(<-cchan)
	//// send data to close chan ,panic
	//cchan<-11

	// 定义只写管道
	var testChan = make(chan int, 60)
	testChan <- 40
	testChan <- 40
	testChan <- 40

	// 定义只读管道 ， select 读管道，不会死锁
	var test2Chan = make(<-chan int)
	select {
	case <-test2Chan:
	default:
		fmt.Println("i do not deadlock")
	}

	// 读写管道
	var test3Chan = make(chan int, 500)
	go func(chan<- int) {
		for i := 0; i < 30; i++ {
			test3Chan <- i
		}
		// 关闭管道
		close(test3Chan)
	}(test3Chan)

	go func(<-chan int) {
		for {
			val, ok := <-test3Chan
			fmt.Println(val)
			if !ok {
				break
			}
		}
	}(test3Chan)

	chanDemo()
	time.Sleep(time.Millisecond)

	for i := 0; i < 3; i++ {
		<-testChan
	}

}

func work(id int) (c chan int) {
	c = make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Printf("helloWord %c -- %d \n", n, id)
		}
	}()

	return c
}

func chanDemo() {
	//语法
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = work(i)
		channels[i] <- 'A' + i

	}

}
