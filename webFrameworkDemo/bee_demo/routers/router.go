package routers

import (
	"bee_demo/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/articles", &controllers.MainController{}, "get:ArticleList")
	beego.CtrlGet("/api/article/get", (*controllers.MainController).ArticleList)
	beego.CtrlPost("/api/article/create", (*controllers.ArticleController).Add)

}
