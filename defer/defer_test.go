package _defer

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	df1()
	fmt.Println("---------------------")
	df2()

}

func df1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func df2() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()

	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()

	}
}

func BenchmarkNodefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithoutDefer()
	}
}

func Sum(Max int) int {
	total := 0
	for i := 0; i < Max; i++ {
		total += i
	}
	return total
}

func fooWithDefer() {
	defer func() {
		Sum(10)
	}()
}

func fooWithoutDefer() {
	Sum(10)
}
