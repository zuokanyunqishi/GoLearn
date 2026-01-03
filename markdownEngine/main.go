package main

import (
	"errors"
	"flag"
	"github.com/88250/lute"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

// markdown 格式化引擎
func main() {

	filePathStr := flag.String("p", "", "文件路径")

	flag.Parse()

	filePath := *filePathStr

	_, err := os.Stat(filePath)

	if errors.Is(err, os.ErrNotExist) {
		log.Fatalf("%s 文件不存在！", filePath)
	}

	engine := lute.New()

	reg := regexp.MustCompile(`^---([\d\D]*)---`)
	if reg == nil {
		log.Fatalf("正则格式错误")
		return
	}

	file, _ := os.Open(filePath)

	defer file.Close()
	all, _ := io.ReadAll(file)
	originStr := string(all)

	result := reg.ReplaceAllStringFunc(originStr, func(s string) string {
		return ""
	})

	builder := strings.Builder{}
	builder.WriteString(reg.FindString(originStr))

	builder.WriteString("\r\n")
	builder.WriteString("\r\n")

	formatStr := engine.FormatStr("demo", result)

	builder.WriteString(formatStr)

	ioutil.WriteFile(filePath, []byte(builder.String()), 0666)

}
