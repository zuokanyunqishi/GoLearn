package main

import (
	"GoLearn/chat/server/route"
	"GoLearn/chat/util"
	"GoLearn/chat/util/zlog"
	"os"
)

// 聊天系统 服务端
func main() {

	path, _ := os.Getwd()

	zlog.Init(path + "/chat/server/storage/server.log")
	server := util.NewServer()
	route.AddRoute(server.(*util.Server))
	server.Serve()
}
