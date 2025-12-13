package controllers

import (
	"errors"
	"net/http"
	"speed/app/lib/validate"
	"time"

	"speed/app/exceptions"
	"speed/app/http/service"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type LoginController struct {
	Controller
	userService *service.UserService
}

func NewLoginController(userService *service.UserService) *LoginController {
	return &LoginController{userService: userService}
}

var LoginC = NewLoginController(service.NewUserService())

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20" comment:"用户名"`
	Password string `json:"password" binding:"required,min=6" comment:"密码"`
}

var tags = make(map[string]string)

func (c *LoginController) Login(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.handleValidationError(ctx, err)
		return
	}

	token, err := c.userService.Authenticate(req.Username, req.Password)
	if err != nil {
		c.handleAuthError(ctx, err, req.Username)
		return
	}

	c.ResponseSuccess(ctx, gin.H{
		"token":   token,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})
}

func (c *LoginController) handleValidationError(ctx *gin.Context, err error) {
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

func (c *LoginController) handleAuthError(ctx *gin.Context, err error, username string) {
	switch {
	case errors.Is(err, exceptions.ErrUserNotFound):
		app.Log.Warn("Login failed: user not found", zap.String("username", username))
		c.ResponseError(ctx, http.StatusUnauthorized, "用户名或密码错误")
	case errors.Is(err, exceptions.ErrWrongPassword):
		app.Log.Warn("Login failed: wrong password", zap.String("username", username))
		c.ResponseError(ctx, http.StatusUnauthorized, "用户名或密码错误")
	default:
		app.Log.Error("Login failed", zap.String("username", username), zap.Error(err))
		c.ResponseError(ctx, http.StatusInternalServerError, "系统错误")
	}
}
