package main

import (
	"embed"
	"io/ioutil"
	app "speed/bootstrap"
	"speed/router"

	"github.com/gin-gonic/gin"
)

//go:embed data/shop.sqlite3
var sqliteDbFile embed.FS

//go:embed .config.json
var configFile embed.FS

func init() {
	app.InitApp(sqliteDbFile, configFile)
}
func main() {

	var ginMode string

	if app.AppEnv == "prod" {
		ginMode = gin.ReleaseMode
		gin.DefaultWriter = ioutil.Discard
	} else {
		ginMode = gin.DebugMode
	}

	gin.SetMode(ginMode)
	engine := gin.Default()
	//engine.LoadHTMLGlob("resources/views/*")
	router.Router(engine) //初始化路由

	_ = engine.Run(":8086")

}
