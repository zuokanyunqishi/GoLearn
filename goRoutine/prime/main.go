package main

import (
	"fmt"
	"time"
)

// 协程求素数

func main() {

	var numStore = make(chan int, 3000)
	var resultStore = make(chan int, 2000)
	var exitStore = make(chan int, 4)

	go putNum(numStore)

	for i := 0; i < 4; i++ {
		go primeHandle(numStore, resultStore, exitStore)
	}

	var exitCount int
	for {
		_, ok := <-exitStore
		if ok {
			exitCount++
		}

		if exitCount >= 4 {
			fmt.Println("工作完成")
			time.Sleep(time.Second * 2)
			break
		}
	}

	for {
		select {
		case value := <-resultStore:
			fmt.Println(value)
		default:
			return
		}
	}

}

func primeHandle(store chan int, store2 chan int, exitChan chan int) {
	for value := range store {

		var flag = true

		for i := 2; i < value/2; i++ {
			if value%i == 0 {
				flag = false
			}
		}

		if flag {
			store2 <- value
		}
	}

	exitChan <- 1

}

func putNum(store chan int) {
	for i := 0; i < 8000; i++ {
		store <- i
	}

	close(store)
}
