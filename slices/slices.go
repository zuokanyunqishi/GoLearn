package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

//切片
//长度 len(s)
//容量 cap(s)

func main() {

	//声明切片方式
	//1 先声明一个数组
	arr := [6]int{8, 2, 255, 32, 45, 0}
	//切片可看做数组的视图
	fmt.Println(arr[:1])  // 取原数组index  1 之前, [8]
	fmt.Println(arr[1:])  // 取原数组 index 1 (含1)之后所有, [2 255 32 45 0]
	fmt.Println(arr[2:5]) // 取原数组 index 2 (含2) 到 index 5 之前 所有, [255 32 45]

	s1 := arr[2:]
	s2 := arr[1:]
	s3 := arr[2:5]

	fmt.Printf("切片s1 = %v ,  length = %d capacity = %d \n", s1, len(s1), cap(s1))
	fmt.Printf("切片s2 = %v , length = %d capacity = %d \n", s2, len(s2), cap(s2))
	fmt.Printf("切片s3 = %v , length = %d capacity = %d \n", s3, len(s3), cap(s3))

	fmt.Println(s2)
	//直接声明切片
	fmt.Println(arr[:])
	fmt.Println(arr[:])
	//update slice
	updateSlice(s1)
	fmt.Println(`After update updateSlice s1`)
	fmt.Println(s1)
	updateSlice(s2)
	fmt.Println(`After update updateSlice s2`)
	fmt.Println(s2)
	updateSlice(s3)
	fmt.Println(`After update updateSlice s3`)
	fmt.Println(s3)

	//slice 是对数组的视窗观察,
	//slice 内部结构
	/*|--------------------
		*|   slice            |
		*|  [ptr] [len] [cap] |
		*---------------------
		*_______________________________
		*| 底层数组                    |
		*|     [][][][][][][][][][][] |
	    *| key 0 1 2 3 4 5 6 7 8 9 10 |
		*|____________________________|
		* ptr 指向slice开始的底层数组位置
		* len 是从ptr开始有有多个元素,长度
		* cap 指向从ptr开始到数组结尾的长度,只能向后扩展
	*/
	sWindow := []int{5, 6, 7, 8, 9, 256}
	sWindow1 := sWindow[1:3]
	sWindow2 := sWindow1[3:5]
	fmt.Println(sWindow1, "\n", sWindow2)
	//打印各个slice的长度的
	fmt.Printf("s_window =%v -- len=%d cap=%d \n", sWindow, len(sWindow), cap(sWindow))
	fmt.Printf("s_window1 =%v -- len=%d cap=%d \n", sWindow1, len(sWindow1), cap(sWindow1))
	fmt.Printf("s_window2 =%v -- len=%d cap=%d \n", sWindow2, len(sWindow2), cap(sWindow2))

	//slice 创建元素,语法 : make(type,len,cap)
	sCreate := make([]int, 6, 42)
	sCreate = fillSlice(sCreate)
	fmt.Println(sCreate)
	//复制一个slice 到另一个slice
	copy(sWindow1, sCreate)
	fmt.Println(sWindow1)
	//删除切片内一个元素
	//用一个切片填充另一个切片
	fmt.Println(append(sCreate[:3], sCreate[6:]...))

	//取切片的第一个,和最后一个

	initSlice := []int{4, 6, 7, 89, 9}
	shift := initSlice[0]
	fmt.Println(`第一个元素是`, shift)
	fmt.Println(`最后一个元素是`, initSlice[len(initSlice)-1])

	repeatStr := strings.Repeat("-", 15)
	fmt.Println(repeatStr + "我是分隔符" + repeatStr)
	fmt.Println("复制切片")

	testSlice := []int{5, 6, 78, 0}

	test2 := make([]int, 1)
	//test 长度只有1 ,testSlice 长度为 4, 只会取testSlice 第1位置
	copy(test2, testSlice)
	fmt.Println(test2)

	fmt.Println(append(test2, testSlice...))

}

func updateSlice(s []int) {

	s[0] = 100
}

func fillSlice(s []int) []int {

	for i := 0; i < 30; i++ {
		s = append(s, 2*i+1)
	}
	return s
}

// base64Decode abc
func base64Decode(str string) string {
	newStr, _ := base64.StdEncoding.DecodeString(str)
	return string(newStr)
}
