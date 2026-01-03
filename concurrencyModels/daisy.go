package concurrencyModels

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

// Whispers 击鼓传花模式
func Whispers() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost

	//fmt.Println(unsafe.Pointer(&right))
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
