package controllers

import (
	"errors"
	"net/http"
	"speed/app/http/model"
	"speed/app/http/service"
	"speed/app/lib/validate"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
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

type EditRequest struct {
	Nickname string `json:"nickname" binding:"omitempty,max=40" comment:"昵称"`
	Sex      int    `json:"sex" binding:"omitempty" comment:"性别"`
	Age      int    `json:"age" binding:"omitempty,min=1,max=150" comment:"年龄"`
	Avatar   string `json:"avatar" binding:"omitempty" comment:"头像"`
	Email    string `json:"email" binding:"omitempty,email,max=50" comment:"邮箱"`
}

func (u *UserController) Edit(ctx *gin.Context) {
	currentUser, err := u.user(ctx)
	if err != nil {
		app.Log.Error("获取用户信息失败", zap.Error(err))
		u.ResponseError(ctx, http.StatusNotFound, "用户未找到")
		return
	}

	var req EditRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		u.handleValidationError(ctx, err)
		return
	}

	updateUser := model.User{
		Nickname: req.Nickname,
		Sex:      req.Sex,
		Age:      req.Age,
		Avatar:   req.Avatar,
		Email:    req.Email,
	}

	err = u.userService.UpdateUser(ctx, currentUser.ID, &updateUser)
	if err != nil {
		app.Log.Error("更新用户失败", zap.Error(err))
		u.ResponseError(ctx, http.StatusInternalServerError, "更新用户失败")
		return
	}

	u.ResponseSuccess(ctx, gin.H{})
}

func (u *UserController) handleValidationError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		app.Log.Error("Invalid request format", zap.Error(err))
		u.ResponseError(ctx, http.StatusBadRequest, "请求格式错误")
		return
	}

	translatedErrs := validate.TranslateError(errs)
	app.Log.Warn("Validation failed", zap.Any("errors", translatedErrs))
	u.ResponseError(ctx, http.StatusBadRequest, translatedErrs)
}
