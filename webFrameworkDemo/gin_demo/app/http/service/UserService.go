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
func (s *UserService) Authenticate(ctx context.Context, username, password string) (*model.User, string, error) {

	var u model.User
	err := u.GetUserByUserName(ctx, username)
	if err != nil {
		return nil, "", exceptions.ErrUserNotFound
	}

	if !hash.CheckPassword(u.Password, password) {
		return nil, "", exceptions.ErrWrongPassword
	}
	generateToken, err := jwt.GenerateToken(u.ID, u.UserName)

	return &u, generateToken, err
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

func (s *UserService) UpdateUser(ctx *gin.Context, userId int, user *model.User) error {
	var u model.User
	err := u.GetUserById(ctx, userId)
	if err != nil {
		return exceptions.ErrUserNotFound
	}

	// 只更新非零值的字段
	if user.Nickname != "" {
		u.Nickname = user.Nickname
	}
	if user.Sex != 0 {
		u.Sex = user.Sex
	}
	if user.Age != 0 {
		u.Age = user.Age
	}
	if user.Avatar != "" {
		u.Avatar = user.Avatar
	}
	if user.Email != "" {
		u.Email = user.Email
	}

	// 确保使用正确的用户ID进行更新
	u.ID = userId
	return u.Update(ctx)
}
