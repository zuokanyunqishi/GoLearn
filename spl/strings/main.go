package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 字符串函数
func main() {
	// 1 . 子串第一次出现的位置
	index := strings.Index("abc", "bc")
	println(index) // 1

	// 2. 子串最后出现位置
	lastIndex := strings.LastIndex("helloword", "d")
	println(lastIndex) // 8

	// 3. 替换子串 n: 替换的个数,-1 替换全部
	replace := strings.Replace("go go hello", "go", "酱油", -1)
	println(replace) // 酱油 酱油 hello

	// 3.1 替换全部子串
	all := strings.ReplaceAll("go go hello", "go", "酱油")
	println(all) //  酱油 酱油 hello

	// 4 以某个子串为标志拆分为数组
	splitArr := strings.Split("php|go|python|java", "|")
	fmt.Printf("splitArr %v \n", splitArr) // splitArr [php go python java]

	// 5. 字符串转大写
	upper := strings.ToUpper("abc bca")
	println(upper) // ABC BCA

	// 6.  字符串转小写
	lower := strings.ToLower("ABC BCA")
	println(lower) // abc bca

	//  7 utf8 字符串遍历
	travelStr := "a small mouse 来 到中国 !"
	travelRune := []rune(travelStr)
	for i := 0; i < len(travelRune); i++ {
		fmt.Printf(" current value is %c \n", travelRune[i])
	}

	// 8. 字符串转整数
	atoi, err := strconv.Atoi("34577")
	if err != nil {
		panic(err)
	}
	println(atoi) // 34577

	// 9. 字符串转整数
	itoa := strconv.Itoa(4568788)
	println(itoa) // "4568788"

	// 10. 字符串转 []byte 数组
	fmt.Printf("%v , \n", []byte("hello word"))

	// 11. 10进制转 2 8 ,16
	println(strconv.FormatInt(8888, 16)) //22b8

	// 12. 子串是否存在在另一个字符串中
	contains := strings.Contains("go lang is best", "best")
	println(contains) // true

	// 13. 查找子串在另一个子串中不重复的count
	count := strings.Count("go go php", "go")
	println(count) // 2

	// 14. 不区分大小写比价字符串
	println(strings.EqualFold("go", "GO")) // true

	// 15. 去除字符串两边空格
	space := strings.TrimSpace(" abc    ")
	fmt.Println(space) // 	abc

	// 16. 去除字符串两边制定字符串
	trim := strings.Trim("!!!!go!!&!", "!&")
	println(trim) // go

	// 17. 去除子串左边
	left := strings.TrimLeft("@@@@php", "@")
	println(left) // php

	// 18. 去除子串右边
	right := strings.TrimRight("Python^^^^^^", "^")
	println(right) // Python

	// 19. 字符串是否以某个子串开头
	prefix := strings.HasPrefix("_env_dev", "_env")
	println(prefix) // true

	// 20. 字符串是否以某个子串结尾
	suffix := strings.HasSuffix("app.jpg", ".jpg")
	print(suffix) // true

}
