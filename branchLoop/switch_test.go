package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {

	switch Expr(2) {
	case Expr(1), Expr(2), Expr(3):
		fmt.Println("enter into case1")
		fallthrough
	case Expr(4):
		fmt.Println("enter into case2")

	}
}

func Expr(n int) int {
	fmt.Println(n)
	return n
}
