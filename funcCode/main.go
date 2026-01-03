package main

import (
	"fmt"
	"time"
)

// 闭包函数
func adder() func(int) int {

	return func(i int) int {
		return i * (i + 1)
	}
}

//传统闭包

type model func(int) int

func (model) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func sum(basic int) model {

	H := basic * 5
	return func(i int) int {
		if basic*H > 9 {
			return 3 * 1
		}

		return i * 1
	}

}

func main() {
	a := adder()

	fmt.Println(a(45))

	fmt.Println(sum(9)(5))

	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, f := range data1 {
		go f.print(3)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}

	for _, f := range data2 {
		f := f
		go (*field).print(&f, 3)
	}

	time.Sleep(time.Second * 3)

}

type field struct {
	name string
}

func (p field) print(i int) {
	fmt.Println(p.name)

}
