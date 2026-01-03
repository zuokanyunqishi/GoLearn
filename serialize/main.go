package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age  int
	Name string
	Work string
}

type PersonTag struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
	Work string `json:"work"`
}

// 序列化反序列化
func main() {
	// 切片
	ints := []int{4, 6, 7, 7}
	intsStr, _ := json.Marshal(ints)
	fmt.Println(string(intsStr))

	// map
	map1 := map[string]interface{}{
		"abc": 124,
		"bda": "abc",
	}
	mapstr, _ := json.Marshal(map1)
	fmt.Println(string(mapstr))

	// 结构体
	person := Person{
		Age:  20,
		Name: "abc",
		Work: "bgh",
	}
	personStr, _ := json.Marshal(person)
	fmt.Println(string(personStr))

	// tag标签序列化
	tag := PersonTag{
		Age:  78,
		Name: "tag",
		Work: "打工人",
	}

	tagStr, _ := json.Marshal(tag)
	fmt.Println(string(tagStr))

}
