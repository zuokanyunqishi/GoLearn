package main

import "fmt"

// 常量
// 关键字 `const`
const filename = `test/file.txt`

// b, kb, mb, gb, tb , pb
const (
	//iota = 0 自增,
	b = 1 << (10 * iota)
	//以下自动套用上面公式,且iota +=1
	kb
	mb
	gb
	tb
	pb
)

const (
	//常量值可以用表达式计算出,函数必须是内置函数,,
	//常量在代码编译时必须可以确定,数字,字符串,布尔值可定义.
	numA = 3 + 4
	numB = `helloWord!`
	numC = len(numB)
	numD = numA * 3
	numF = true
)

const (
	//声明了 iota 底层就一直累加
	aa = iota //0
	bb        //2
	cc
	d = "ha" //独立值，iota += 1
	e        //"ha"   iota += 1
	f = 100  //iota +=1
	g        //100  iota +=1
	h = iota //7,恢复计数
	i        //8
	_        //跳过 iota +=1
	j        //10
)

func main() {

	//函数内常量
	//多重赋值
	const length, width = 3, 4

	fmt.Println(b)
	fmt.Println(kb)
	fmt.Println(mb)
	fmt.Println(tb)
	fmt.Println(gb)
	fmt.Println(pb)
	fmt.Println(filename)
	fmt.Printf("矩形面积是 %d 平方 \n", length*width)
	fmt.Println(numA, numB, numC, numD, numF)
	fmt.Println(aa, bb, cc, d, e, f, g, h, i, j)
}
