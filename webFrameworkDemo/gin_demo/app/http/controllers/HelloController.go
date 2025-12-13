package controllers

import (
	"speed/app/lib/log"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
	Controller
}

var HelloC = &HelloController{}

func (h *HelloController) Index(ctx *gin.Context) {

	// fmt.Println(model.Users{}.GetMore())
	log.WithCtx(ctx).Info("hello word")
	h.ResponseSuccess(ctx, map[string]interface{}{"hello": "hello word"})
}
