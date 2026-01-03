package middleware

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ExceptionHandle(ctx context.Context) ghttp.HandlerFunc {

	return func(r *ghttp.Request) {
		r.Middleware.Next()
		if err := r.GetError(); err != nil {

			// 记录到自定义错误日志文件
			g.Log("exception").Error(ctx, err)
			//返回固定的友好信息
			r.Response.ClearBuffer()
			r.Response.Writeln("服务器居然开小差了，请稍后再试吧！")
		}
	}

}
