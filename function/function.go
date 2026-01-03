package main

import (
	"fmt"
)

//函数
//go函数支持多返回值,返回值支持变量名
//支持可变参数
//支持闭包,匿名函数
//多返回值可忽略接收某一个

func main() {

	fmt.Println(demo(78, 56))
	fmt.Println(sum(3, 4, 5, 7))

	//多返回值可忽略某一个接收, `_`表示忽略
	a, _ := moreReturn(125, 390)
	fmt.Println(a)

	//匿名函数作为参数使用
	result := functionParam(func(a, b int) int {
		return a * b
	}, 5, 6)

	fmt.Println(result)

	//直接传入本包内的 函数名称
	resultA := functionParam(functionParamA, 3, 4)

	fmt.Println(resultA)

	//闭包
	closure1, closure2 := closure(), closure()

	closure1() //i=1,内存地址0xc42007e040
	closure1() //i=2,内存地址0xc42007e040

	closure2() //i=1,内存地址0xc42007e048
	closure2() //i=2,内存地址0xc42007e048
	closure2() //i=3,内存地址0xc42007e048
}

func demo(paramA, paramB int) interface{} {
	return paramA + paramB

}

// 可变参数
func sum(number ...int) int {
	result := 0

	for i := range number {

		result += i
	}
	return result

}

// 多返回值 , 返回值变量名
func moreReturn(a, b int) (value1, value2 int) {

	value1 = a * b
	value2 = a + b
	//返回值声明变量,可直接return,底层会自动寻找返回,但不建议这麽做
	return

}

// 函数当成参数,参数op为函数,可指定该函数的具体规则
func functionParam(op func(a, b int) int, x, y int) int {

	return op(x, y) + functionParamA(5, 6)
}

// 定义一个函数,直接传入函数名
func functionParamA(w, x int) int {

	return w*w + x*x
}

// 闭包函数
func closure() func() {
	i := 0
	return func() {
		//包内变量i是对函数包外的变量的引用,--内存地址
		i++
		fmt.Printf("变量--i-- 值是 %d,内存地址是 [%p]\n", i, &i)
	}
}
