package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// FileNotFound 自定义错误
var FileNotFound = errors.New("文件不存在")
var TimeOver = errors.New("时间结束")

func main() {

	err := readFile()
	if errors.Is(err, FileNotFound) {
		fmt.Println(err)
	}

	after := time.After(time.Second * 5)

	select {
	case <-after:
		fmt.Println(TimeOver)

	}

}

func readFile() error {
	if _, err := os.Open("xxxx"); err != nil {
		return FileNotFound
	}
	return nil
}
