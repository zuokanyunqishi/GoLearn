package main

import (
	"GoLearn/testTool"
	"fmt"
)

// map
func main() {
	//声明 map
	//语法 map[type]type{}
	var smallMap = map[string]int{
		`a`: 12,
		`b`: 13,
		`c`: 56,
	}

	fmt.Println(smallMap)

	//make 语法
	bigMap := make(map[int]string)
	bigMap[2] = `helloWord`
	bigMap[3] = `helloWord`
	bigMap[4] = `helloWord`
	bigMap[5] = `helloWord`
	testTool.Dd(bigMap)

	//遍历map

	for key, value := range bigMap {
		fmt.Println(`key `, key, `value`, value)
	}

	//删除map中元素
	delete(bigMap, 2)
	fmt.Println(bigMap)
	//寻找最长不含有重复字符串的子串

	fmt.Println(lengthOfLongestSubstring(`abcdefghf`))
	fmt.Println(lengthOfLongestSubstring(`wazw`))
	fmt.Println(lengthOfLongestSubstring(`opytq`))
	fmt.Println(lengthOfLongestSubstring(`aaa`))
	fmt.Println(lengthOfLongestSubstring(``))
	fmt.Println(lengthOfLongestSubstring(`pppo`))
}

func lengthOfLongestSubstring(s string) int {

	lastRepeat := make(map[byte]int)
	start, maxLength := 0, 0
	for i, ch := range []byte(s) {
		lastI, ok := lastRepeat[ch]
		if lastI >= start && ok {
			start = lastI + 1
		}

		if maxi := i - start + 1; maxi > maxLength {
			maxLength = maxi
		}

		lastRepeat[ch] = i
	}
	return maxLength
}
