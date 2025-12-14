package middleware

import (
	"net/http"
	"speed/app/http/model"
	"speed/app/lib/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			ResponseError(c, "未提供Token")
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ") // 去除 "Bearer "

		// 2. 验证 Token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			ResponseError(c, "无效Token")
			c.Abort()
			return
		}
		// 3. 将用户信息存入上下文
		c.Set("userId", claims.UserID)
		User := model.User{}
		_ = User.GetUserById(c, claims.UserID)
		c.Set("user", User)
		c.Next()
	}
}

func ResponseError(ctx *gin.Context, message interface{}) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code":    http.StatusUnauthorized,
		"message": message,
		"data":    gin.H{},
	})
}
