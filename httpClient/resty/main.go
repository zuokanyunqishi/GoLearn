package main

import (
	"fmt"
)

func main() {

	//client := resty.New()
	//
	//response, _ := client.R().Get("https://api.bilibili.com/x/web-interface/search/default")
	//
	//var json interface{}
	//sonic.Unmarshal(response.Body(), &json)
	//testTool.Dd(json)
	A()

}

func A() {
	c := make(chan int)
	ok := make(chan struct{})

	go func() {
		for item := range c {
			fmt.Println(item)
		}
		ok <- struct{}{}
	}()

	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
	<-ok
}
