package main

import (
	"GoLearn/errorHanding/tryCatch"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	. "net/http"
	"os"
)

func tryCatch() {
	fmt.Println(1)

	//defer 先入后出栈,3先显示,然后是2
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println(4)
}

func createFile(fileName string) {
	writer, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	//关闭资源
	defer writer.Close()

	//写入buffer缓存
	buffer := bufio.NewWriter(writer)

	//将缓存写入文件
	defer buffer.Flush()

	for i := 0; i < 40; i++ {
		fmt.Fprintln(writer, i*3)
	}
}

func loopAndPanic() {
	for i := 0; i < 50; i++ {
		defer fmt.Println(i)
		if i > 40 {
			panic(`禁止循环`)
		}

	}
}

type attr struct {
	fun func(w ResponseWriter, r *Request)
}

func main() {
	//tryCatch()
	//
	//createFile(`aaa.txt`)
	//loopAndPanic()

	HandleFunc("/list/helloword.txt",
		response)

	log.Fatal(ListenAndServe(":8080", nil).Error())

}

func response(w ResponseWriter, r *Request) {
	path := r.URL.Path
	bytes, er := ioutil.ReadFile(path)
	if er != nil {
		fmt.Println(er.Error())
		httpError.ErrorHandle(w, er)
	}
	w.Write(bytes)

}
