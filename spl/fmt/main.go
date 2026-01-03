package main

import "fmt"

func main() {

	var arr [6]float64
	// 1. 从终端读取
	for i := 0; i < len(arr); i++ {
		fmt.Printf("请输入第 %d个数字\n", i+1)
		scanln, _ := fmt.Scanln(&arr[i])
		fmt.Println(scanln)
	}

}
