package main

import (
	"fmt"
	"unicode/utf8"
)

//中文字符处理

func main() {
	s := `yes我是中国人!`
	printChina(s)
	//打印字节
	for i, b := range `我爱中国` { //b is rune
		fmt.Printf(`(%d) %X %v`, i, b, ` `)

	}
	//utf8 库处理中文
	fmt.Println("\n", `中文字长`, utf8.RuneCountInString(`你好世界`))

	//转化为rune数组
	for i, ch := range []rune(s) {
		fmt.Printf("%d %X %c \n", i, ch, ch)
	}

	//字符串分割 split,join
	//Trim
	//ToLower
	//Index,Contains

	// 字串串底层是一个只读的 []byte 切片,
	// 要修改字符串先转为切片修改,再转成字符串.
	// 带有中文字符的转为 []rune 切片
	str := "hello,china"
	sar := []byte(str)
	sar[0] = '!'
	println(string(sar))
}

func printChina(s string) {
	for i, b := range s {
		fmt.Println(b, i)
		fmt.Printf(" %d%c \n", i, b)

	}

}
