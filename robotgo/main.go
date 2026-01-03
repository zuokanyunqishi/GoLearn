package main

import (
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-vgo/robotgo"
	"github.com/go-vgo/robotgo/clipboard"
	"net/http"
	"os"
)

func main() {

	_, err2 := os.StartProcess("/usr/bin/google-chrome-stable", []string{"%U"}, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if err2 != nil {
		return
	}

	robotgo.Sleep(5)

	robotgo.MoveSmooth(140, 90)
	robotgo.TypeStr("url")
	robotgo.KeyTap("enter")

	robotgo.Sleep(3)

	robotgo.KeyTap("r", "ctrl")
	robotgo.Sleep(3)
	robotgo.KeyTap("s")
	robotgo.KeyTap("r")

	robotgo.Sleep(5)

	robotgo.KeyTap("c")
	robotgo.KeyTap("p")
	robotgo.Sleep(5)

	text, err := clipboard.ReadAll()

	json := make(map[string]string)
	json["key"] = "xxx"
	json["title"] = "test"
	json["public"] = "0"
	json["format"] = "markdown"
	json["body"] = text

	marshal, _ := sonic.Marshal(json)
	fmt.Println(string(marshal))
	url := "https://www.yuque.com/api/v2/repos/xxxxx/docs"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshal))
	req.Header.Set("X-Auth-Token", "xxxxxx")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err2 := client.Do(req)
	fmt.Println(err, response.Body.Close())

}
