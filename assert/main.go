package main

import "fmt"

// 类型断言

func main() {

	var a interface{}

	b := 60
	a = b
	fmt.Println(a)
	var c int
	// 空接口不可直接赋值
	c = a.(int)
	fmt.Println(c)

	typeJudge([]interface{}{1, nil, struct {
	}{}, 13.45, "str"}...) // 可变参数

}

func typeJudge(item ...interface{}) {

	for _, value := range item {

		switch value.(type) {
		case int:
			fmt.Println("value is int")
		case nil:
			fmt.Printf("value is nil %T\n", value)
		case float32, float64:
			fmt.Printf("value is float %T\n", value)
		case struct{}:
			fmt.Printf("value is struct %T\n", value)
		case string:
			fmt.Printf("value is string %T\n", value)

		}
		fmt.Println()

	}

}
