package main

import (
	"fmt"
	"strings"
)

// 数组
// 是一种线性表的数据结构.用一组连续的内存空间,存储一组具有相同类型的数据
// go 的数组 `[5]int{..}` `[...]string{...}`
func main() {

	/*声明数组方法*/
	// 1. 指定长度大小
	var arr = [5]int{1, 2, 3, 4, 5}
	//初始化数组
	var arrInt [3]int         // 整数型初始化 填充0 [0,0,0]
	var arrStr [3]string      // 字符串类型初始化 填充0 ['','','']
	var arrStruct [3]struct{} //结构体初始化 [{},{},{}]
	fmt.Println(arrInt, arr, arrStr, arrStruct)

	//2 .省略长度 让编译器计算长度
	arrAuto := [...]int{2, 11, 13, 14}
	fmt.Println(`数组是`, arrAuto, `长度是`, len(arrAuto))

	/* 指定数组某个索引的值 */
	arrIndex := [4]int{1: 2, 3: 1} //[0 2 0 1]
	str := `Array
	 -----------------------------------
        | 0 | 1 | 2 | 3 | 4 | index
	 ---|---|---|---|---|---|-----------
        | 0 | 2 | 0 | 1 | 0 | value
	 ---|---|---|---|---|---|-----------
	`
	fmt.Println(`指定index的值的array `, arrIndex)
	fmt.Printf("结构如下 %s  \n", strings.Repeat(`+`, 32))
	fmt.Println(str)

	//遍历数组
	traversalArr(arrAuto)
	//函数传递数组指针
	arrP := [4]int{}
	fmt.Println(arrP)
	arrayPointer(&arrP)
	fmt.Println(arrP)

	//定义多维数组

	var arr2 = [3][3][3]int{{{1, 2, 3}}}
	fmt.Println(arr2)

	//翻转数组
	fmt.Println(rotateArr([]int{4, 34, 78, 90, 8}))
}

// 遍历数组
func traversalArr(arr [4]int) {
	//常规遍历方法
	for i := 0; i < len(arr); i++ {
		fmt.Printf("index [%d]  value [%d] \n", i, arr[i])
	}

	// for range 方法
	for i, v := range arr {
		fmt.Println(`index`, i, `value`, v)
	}

	//省略 index
	for _, v := range arr {
		fmt.Println(`value`, v)
	}
}

// 数组指针
func arrayPointer(arrP *[4]int) {

	arrP[1] = 4
	fmt.Println(arrP)
	//index 1 地址
	fmt.Println(&arrP[1])
	fmt.Println(&arrP[2])
	fmt.Println(&arrP[3])
}

// 翻转数组
func rotateArr(arr []int) []int {

	tmp := 0
	for i := 0; i < len(arr)/2; i++ {
		tmp = arr[len(arr)-1-i]

		arr[len(arr)-1-i] = arr[i]
		arr[i] = tmp
	}

	return arr
}
