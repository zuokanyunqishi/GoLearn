package middleware

import (
	"net/http"
	"speed/app/lib/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 获取 Token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供Token"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ") // 去除 "Bearer "

		// 2. 验证 Token
		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效Token"})
			return
		}

		// 3. 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
