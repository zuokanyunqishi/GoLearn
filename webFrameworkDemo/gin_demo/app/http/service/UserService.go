package service

import (
	"context"
	"speed/app/exceptions"
	"speed/app/http/model"
	"speed/app/lib/hash"
	"speed/app/lib/jwt"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

// Authenticate Service层实现
func (s *UserService) Authenticate(ctx context.Context, username, password string) (string, error) {

	var u model.User
	err := u.GetUserByUserName(ctx, username)
	if err != nil {
		return "", exceptions.ErrUserNotFound
	}

	if !hash.CheckPassword(u.Password, password) {
		return "", exceptions.ErrWrongPassword
	}
	generateToken, err := jwt.GenerateToken(u.ID, u.UserName)

	return generateToken, err
}

func (s *UserService) AddUser(ctx *gin.Context, user *model.User) error {
	db_u := model.User{}
	err := db_u.GetUserByUserName(ctx, user.UserName)
	if err == nil && db_u.UserName == user.UserName {
		return exceptions.UserNameIsUsed
	}

	return user.Add(ctx)
}

func (s *UserService) Me(ctx *gin.Context, userId int) (model.User, error) {
	var u model.User
	err := u.GetUserById(ctx, userId)
	if err != nil {
		return u, err
	}
	return u, nil
}
