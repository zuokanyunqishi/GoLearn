package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*内建标准变量类型*/
func main() {

	//bool
	isTrue := false

	//string 用 `` 和 ""
	str := `helloWord!`
	str1 := "helloChina"

	/*******数字类型*****/
	//无符号整数
	var numUint8 uint8 = 255                    //无符号整数,取值范围[0 -- 255]
	var numUint16 uint16 = 65535                //无符号整数,取值范围[0 -- 65535]
	var numUint32 uint32 = 4294967295           //无符号整数,取值范围[0 -- 42 9496 7295]
	var numUint64 uint64 = 18446744073709551615 //无符号整数,取值范围[0 -- 1844 6744 0737 0955 1615]

	//有符号整数数
	var numInt8 int8 = -128                   //取值范围 -128 --- 127
	var numInt16 int16 = -32768               //取值范围 -32768 ---32767
	var numInt32 int32 = -2147483648          //取值范围 -2147483648 --- 2147483647
	var numInt64 int64 = -9223372036854775808 //取值范围 -9223372036854775808 --- 9223372036854775807

	//浮点型
	var numFloat32 float32 = float32(numUint64)
	var numFloat64 float64 = float64(numUint64)

	//实数虚数 todo
	//var numComplex64  complex64 = 3 +  4i
	//var numcomplex128  complex128 =

	/***** 其他数字类型  */

	//byte
	var numByte byte = 255 //类似 uint8 ,取值范围 [0 -- 255] 16进制,八进制
	//rune
	var numRune rune = -2147483648 //类似 有符号数 int32 ;取值范围 [-2147483648 --- 2147483647]
	//uint
	var numUint uint = 18446744073709551615 //(无符号数)类似 uint32 或者 uint64 ,根据操作系统是32位 || 64 位
	//int
	var numInt int = -9223372036854775808 //(有符号数)类似int32 或者 int64 ,取值范围根据操作系统32位 || 64位
	//uintptr 指针  todo

	//var numPoint uintptr = 11222222  //(无符号整型)

	fmt.Println(`布尔值 ---`, isTrue)
	fmt.Println(`字符串 ---`, str)
	fmt.Println(`字符串 ---`, str1)

	fmt.Println(strings.Repeat(`-`, 40))

	fmt.Println(`uint8 ---`, numUint8)
	fmt.Println(`uint16 ---`, numUint16)
	fmt.Println(`uint32 ---`, numUint32)
	fmt.Println(`uint64 ---`, numUint64)

	fmt.Println(strings.Repeat(`-`, 40))

	fmt.Println(`int8 ---`, numInt8)
	fmt.Println(`int16 ---`, numInt16)
	fmt.Println(`int32 ---`, numInt32)
	fmt.Println(`int64 ---`, numInt64)

	fmt.Println(strings.Repeat(`-`, 40))

	fmt.Println(`float32 ---`, numFloat32)
	fmt.Println(`float64 ---`, numFloat64)

	fmt.Println(strings.Repeat(`-`, 40))

	//fmt.Println(`complex64 ---` ,numComplex64)

	fmt.Println(`byte ---`, numByte)
	fmt.Println(`rune ---`, numRune)
	fmt.Println(`uint ---`, numUint)
	fmt.Println(`int ---`, numInt)
	//fmt.Println(`uintptr ---`,numPoint)

	fmt.Println(strings.Repeat(`-`, 40))

	//强制类型转换
	typeConversion()
}

// 强制类型转换
func typeConversion() {
	//数字字符串转数字
	numStr := `123456`
	num, _ := strconv.Atoi(numStr)

	var a, b int = 3, 6
	var c int

	c = a*a + b*b
	//强制类型转换位float64
	d := math.Sqrt(float64(c))
	//c 声明为int , <math.Sqrt> 返回值是 float64 ,必须强制转换位int
	c = int(math.Sqrt(float64(c)))

	fmt.Println(num)
	fmt.Println("int --", c)
	fmt.Println("float64 --", d)
}
