package middleware

import (
	"speed/app/lib"
	"speed/app/lib/log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Trace(ctx *gin.Context) {
	ctx.Set("traceId", uuid.New().String())
	uri := ctx.Request.URL.String()
	ctx.Set("uri", uri)
	ctx.Set("startTime", lib.MsTime())
	log.WithCtx(ctx).Info("service start")

	ctx.Next()
	log.WithCtx(ctx).Infof("service end --- 耗时 %v ms", lib.MsTime()-ctx.Value("startTime").(int64))

}
