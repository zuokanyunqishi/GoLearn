package main

import (
	"flag"
	"fmt"
	"os"
)

// 命令行参数解析
func main() {

	// 终端参数列表 切片
	args := os.Args
	fmt.Println(args)

	var age int
	var name string
	flag.IntVar(&age, "a", 10, "年龄")
	flag.StringVar(&name, "n", "", "姓名")

	flag.Parse()

	fmt.Println(age, name)

}
