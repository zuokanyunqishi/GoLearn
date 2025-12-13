package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// 统一响应格式封装
func (*Controller) ResponseError(ctx *gin.Context, status int, message interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    status,
		"message": message,
		"data":    gin.H{},
	})
}

func (*Controller) ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
}

// TextRes f 格式化的字符串
// s 输出的字符串
func (*Controller) TextRes(c *gin.Context, f string, s ...string) {
	c.String(http.StatusOK, f, s)
}
