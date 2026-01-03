package main

import "fmt"

//变量定义,用 `var` 或 `:=` ,

//包内定义变量,也就是函数外,只可用 `var`

var (
	//批量定义
	test1        = `hello`
	test2        = 13
	test3        = true
	test4 uint32 = 123
)

func main() {

	//函数内作用域
	var num1 int = 1
	var str string = `string`
	var boole bool
	//`:=` 只能在函数内使用
	//自动类推断
	a, b, c := `j`, 23, 12.5
	fmt.Println(num1, str, boole)
	fmt.Println(a, b, c)
	fmt.Println(test1, test2, test3, test4)
}
