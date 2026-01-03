package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		errorRecover()

		if i > 45 {
			break
		}
		i++
	}
}

func errorRecover() {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)

		}
	}()

	b := 0
	a := 5 / b
	fmt.Println(a)
}
