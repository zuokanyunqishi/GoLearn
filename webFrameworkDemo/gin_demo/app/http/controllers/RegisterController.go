package controllers

import (
	"errors"
	"net/http"
	"speed/app/exceptions"
	"speed/app/http/model"
	"speed/app/http/service"
	"speed/app/lib/hash"
	"speed/app/lib/validate"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type RegisterController struct {
	Controller
	userService *service.UserService
}

func New(userService *service.UserService) *RegisterController {
	return &RegisterController{userService: userService}
}

var Register = New(service.NewUserService())

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20" comment:"用户名"`
	Password string `json:"password" binding:"required,min=6" comment:"密码"`
	Phone    string `json:"phone" binding:"required,min=11,max=11" comment:"手机号"`
}

func (c *RegisterController) Register(ctx *gin.Context) {
	var req RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.handleValidationError(ctx, err)
		return
	}
	hashPwd, _ := hash.HashPassword(req.Password)
	u := model.User{Password: hashPwd, UserName: req.Username, Phone: req.Phone}
	err := c.userService.AddUser(ctx, &u)
	if err != nil {
		c.handleAuthError(ctx, err, req.Username)
		return
	}
	c.ResponseSuccess(ctx, gin.H{})
}

func (c *RegisterController) handleValidationError(ctx *gin.Context, err error) {
	var errs validator.ValidationErrors
	ok := errors.As(err, &errs)
	if !ok {
		app.Log.Error("Invalid request format", zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, "请求格式错误")
		return
	}

	translatedErrs := validate.TranslateError(errs)
	app.Log.Warn("Validation failed", zap.Any("errors", translatedErrs))
	c.ResponseError(ctx, http.StatusBadRequest, translatedErrs)
}

func (c *RegisterController) handleAuthError(ctx *gin.Context, err error, username string) {
	switch {
	case errors.Is(err, exceptions.UserNameIsUsed):
		c.ResponseError(ctx, http.StatusBadRequest, err.Error())
	default:
		app.Log.Error("Login failed", zap.String("username", username), zap.Error(err))
		c.ResponseError(ctx, http.StatusBadRequest, "系统错误")
	}
}
