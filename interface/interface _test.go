package main

import (
	"fmt"
	"reflect"
	"testing"
)

func DumpMethod(i interface{}) {
	v := reflect.TypeOf(i)
	elem := v.Elem()
	n := elem.NumMethod()
	if n == 0 {
		fmt.Println("method is empty")
		return
	}

	fmt.Println("method for this")

	for j := 0; j < n; j++ {
		fmt.Println("-", elem.Method(j).Name)
	}
}

type Interface interface {
	M11()
	M22()
}

type T struct{}

func (t T) M11() {

}

func (t *T) M22() {

}

func TestInterface(test *testing.T) {
	var t T
	var t1 *T

	var i Interface

	i = t1
	// i = t  cannot use t (variable of type T) as type Interface in assignment: T does not implement Interface (M22 method has pointer receiver)

	fmt.Println(i)
	DumpMethod(&t)
	DumpMethod(&t1)
	DumpMethod((*Interface)(nil))

}
