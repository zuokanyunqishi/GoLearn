package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 文件操作
func main() {

	// 文件操作
	// 缓冲方式打开文件
	workPath, _ := os.Getwd()
	filePath := workPath + "/file/text.txt"
	file, _ := os.Open(filePath)
	reader := bufio.NewReader(file)
	readStr, _ := reader.ReadString('\n')
	file.Close()
	fmt.Println("缓冲方式读写", readStr)

	// 直接全部读取文件
	fileContent, _ := ioutil.ReadFile(filePath)

	// 读的是字节, 转化成字符串
	fmt.Println("直接全部方式读写", string(fileContent))

	// 以追加写入创建文件

	openFile, err := os.OpenFile(workPath+"/file/append.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error ", err)
	}

	// 以 缓冲写入方式

	write := bufio.NewWriter(openFile)

	for i := 0; i < 100; i++ {
		write.WriteString("打酱油,打酱油....\n")
	}

	write.Flush()

}
