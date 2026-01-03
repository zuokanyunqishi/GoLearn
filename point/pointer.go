package main

import (
	"fmt"
	"strings"
)

//指针
//指针可以理解为指针变量,变量在内存中的地址0X... 付给了另外一个变量
//`&` 取变量的内存地址
//go 函数调用传参数是值传递,复制

func main() {
	var v = `helloWord`
	//`&` 取内存地址并赋值给另外一个变量
	var point = &v

	fmt.Printf("指针变量point = %p  对应的值是 %s \n", point, *point)
	//通过指针变量存储的 `v` 的内存地址,找到v的值并改变
	*point = `helloChina`
	//打印指针变量所指向的值
	fmt.Println(`改变变量v的值 `, *point)

	//传值的函数调用
	var param = 100
	fmt.Println(`没有调用 变量param的内存地址 `, &param) //0xc42007e010
	passByVal(param)                          //0xc42007e018

	fmt.Println(strings.Repeat(`-`, 64))

	//传递指针的函数调用
	var ref = `pass_ref`
	fmt.Println(`没有调用 变量ref的内存地址是--`, &ref)
	fmt.Println(`没有调用 变量ref的值是--`, ref)
	passByRef(&ref)

}

// 函数调用,传递值
func passByVal(param int) {
	fmt.Println(`调用函数 变量param的内存地址 `, &param)
}

// 传递指针
func passByRef(param *string) {
	//更改指针所指的值
	*param = `ha ha ha`
	fmt.Println(`调用函数 传递的指针变量值是--`, param)
	fmt.Println(`调用函数 传递的指针变量值的 内存地址 是--`, &param)
	fmt.Println(`调用函数 更改后的变量 ref 值的内存地址是--`, &*param)
	fmt.Println(`调用函数 更改后的变量 ref 值是--`, *param)
}
