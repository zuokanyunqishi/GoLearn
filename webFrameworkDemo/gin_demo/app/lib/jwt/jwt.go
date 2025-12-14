package jwt

import (
	"fmt"
	"speed/app/exceptions"
	app "speed/bootstrap"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * 10 * time.Hour)), // 过期时间
			Issuer:    "wawa_shop",                                             // 签发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(app.AppKey)) // 密钥需保密！
}

type CustomClaims struct {
	UserID               int    `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 嵌入标准声明（ExpiresAt, Issuer等）
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 1. 验证签名算法（防算法替换攻击）
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(app.AppKey), nil
		},
	)
	if err != nil {
		return nil, err // 解析失败（如格式错误）
	}

	// 2. 提取 Claims 并验证有效性
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 3. 验证标准声明（过期时间、生效时间等）
		if time.Now().After(claims.ExpiresAt.Time) {
			return nil, exceptions.JwtTokenExpire
		}
		return claims, nil
	}
	return nil, exceptions.JwtTokenInvalid
}
