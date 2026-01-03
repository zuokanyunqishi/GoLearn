package main

import (
	"GoLearn/chat/client/route"
	"GoLearn/chat/client/service"
	"GoLearn/chat/util/zlog"
	"os"
)

func init() {
	path, _ := os.Getwd()
	zlog.Init(path + "/chat/client/storage/client.log")
}
func main() {

	client := service.NewClient()

	route.AddRoute(client.GetServer())

	client.Run()

}
