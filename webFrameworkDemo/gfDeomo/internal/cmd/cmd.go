package cmd

import (
	"context"
	"gfDeomo/internal/middleware"
	"gfDeomo/utility/response"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gfDeomo/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {

				group.Middleware(middleware.ExceptionHandle(ctx), func(r *ghttp.Request) {

					r.Middleware.Next()

					// 如果已经有返回内容，那么该中间件什么也不做
					if r.Response.BufferLength() > 0 {
						return
					}

					var (
						err             = r.GetError()
						res             = r.GetHandlerResponse()
						code gcode.Code = gcode.CodeOK
					)
					if err != nil {
						code = gerror.Code(err)
						if code == gcode.CodeNil {
							code = gcode.CodeInternalError
						}
						response.JsonExit(r, code.Code(), err.Error())
						//if r.IsAjaxRequest() {
						//	response.JsonExit(r, code.Code(), err.Error())
						//} else {
						//	service.View().Render500(r.Context(), model.View{
						//		Error: err.Error(),
						//	})
						//}
					} else {
						response.JsonExit(r, code.Code(), "", res)
						//if r.IsAjaxRequest() {
						//	response.JsonExit(r, code.Code(), "", res)
						//} else {
						//	// 什么都不做，业务API自行处理模板渲染的成功逻辑。
						//}
					}
				})

				group.Bind(
					controller.Hello,
					controller.Article)
			})

			s.EnableAdmin()
			s.Run()
			return nil
		},
	}
)
