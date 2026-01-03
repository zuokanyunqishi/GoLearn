package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

const file_name = `../test/file.txt`

//流程控制 和 循环

func main() {

	//branchIf()

	//fmt.Println(branchSwitch(2, 3, `*`))
	//fmt.Println(branchSwitch2(6))

	//二进制,do while
	fmt.Println(loopFor(1000))
	//模仿while循环
	loopFor1()
	//死循环
	//loopFor2()
	//正常for
	fmt.Println(loopForNormal())
	//Block()

	if1()
}

// if流程控制
func branchIf() {
	//读取文件
	contents, err := ioutil.ReadFile(file_name)
	//if 一般写法
	if err == nil {
		fmt.Printf("%s", contents)

	} else {
		fmt.Println(err)
	}

	//if 可以向下边这样写,if语句快内的--变量仅在块内有效--,有变量作用域限制
	if content, errInfo := ioutil.ReadFile(file_name); errInfo != nil {
		fmt.Println(errInfo)
	} else {
		fmt.Printf("%s \n ", content)
	}

}

// switch 分之 ,和其他语言的switch 分之不同,没有break
func branchSwitch(a, b int, op string) int {
	//result := 0;
	var result int
	switch op {
	case `+`:
		result = a + b
	case `-`:
		result = a - b
	case `*`:
		result = a * b
	case `/`:
		result = a / b
	default:
		result = 0
	}
	return result
}

// switch 另外一种写法
func branchSwitch2(score int) string {

	if score < 0 || score > 100 {
		//调试用,报错终止程序执行
		panic("成绩不在判断范围!")
	}

	msg := ``
	switch {
	case score > 90:
		msg = `A`
	case score > 70:
		msg = `B`
	case score > 60:
		msg = `C`
	default:
		msg = `H`
	}
	return `成绩等级是 --- ` + msg
}

// for循环 10进制转2机制
func loopFor(n int) string {
	//for循环
	binary := ``
	// do while
	for ; n > 0; n /= 2 {

		low := n % 2
		//强制转字符串
		binary = strconv.Itoa(low) + binary
	}

	return binary
}

// while 循环
func loopFor1() {
	//while 循环
	n := 0
	for n < 10 {
		n++
		fmt.Println(`hello +`, n)
	}
}

// 死循环
func loopFor2() {
	for {
		fmt.Println(`helloWord`)
	}
}

// 正常循环
func loopForNormal() int {

	sum := 0
	for i := 1; i < 100; i++ {
		//计算奇数的和
		if i%2 > 0 {
			continue
		}

		sum += i
		if sum > 15 {
			break
		}
	}

	return sum
}

// Block 作用域
//func Block() {
//	if i := 1; true {
//		i++
//		var j = 3
//		j++
//	} else {
//		fmt.Println(i, j)
//	}
//
//	fmt.Println(j)
//}

//func for1() {
//
//	{
//
//		for j, k := 0, 5; j < 10; j++ {
//			fmt.Println(j, k)
//		}
//	}
//	fmt.Println(j)
//
//}

func if1() {
	if a := 1; false {

	} else if b := 2; false {

	} else if c := 3; false {

	} else {
		fmt.Println(a, b, c)
	}
}
