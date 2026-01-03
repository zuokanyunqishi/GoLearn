package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.Default()
	app.Get("/", func(context context.Context) {
		context.Path()
		context.JSON(map[string]string{"a": "c"})
	})

	app.Listen(":8080")
}
