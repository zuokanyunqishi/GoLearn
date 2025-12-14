package controllers

import (
	"net/http"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"

	"speed/app/http/service"
)

type UserController struct {
	Controller
	userService *service.UserService
}

var UserC = &UserController{userService: &service.UserService{}}

func (u *UserController) Me(ctx *gin.Context) {
	id, exists := ctx.Get("userId")
	if !exists {
		u.ResponseError(ctx, http.StatusNotFound, "用户未找到")
		return
	}
	me, err := u.userService.Me(ctx, id.(int))
	if err != nil {
		u.ResponseError(ctx, http.StatusNotFound, "用户未找到")
	}
	u.ResponseSuccess(ctx, me)
}

func (u *UserController) Edit(ctx *gin.Context) {
	user, err := u.user(ctx)
	if err != nil {
		app.Log.Error(err)
		return
	}
	app.Log.Info(user)
}
