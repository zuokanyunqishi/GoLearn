package model

import (
	"context"
	app "speed/bootstrap"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID            int    `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	UserName      string `gorm:"type:text;size:30;not null;column:username" json:"username"`
	Phone         string `gorm:"type:text;size:11;not null;column:phone" json:"phone"`
	Password      string `gorm:"type:text;size:255;not null;column:password" json:"-"`
	Nickname      string `gorm:"type:text;size:40;column:nickname" json:"nickname"`
	Sex           int    `gorm:"type:integer;column:sex" json:"sex"`
	Age           int    `gorm:"type:integer;column:age" json:"age"`
	Avatar        string `gorm:"type:text;column:avatar" json:"avatar"`
	Email         string `gorm:"type:text;size:50;column:email" json:"email"`
	LastLoginTime int64  `gorm:"type:integer;not null;default:0;column:last_login_time" json:"last_login_time"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserByUserName(ctx context.Context, username string) error {
	tx := app.Db.WithContext(ctx).Where("username = ?", username).Find(&u)
	return tx.Error
}

func (u *User) Add(ctx context.Context) error {
	tx := app.Db.WithContext(ctx).Save(&u)
	return tx.Error
}

func (u *User) GetUserById(ctx *gin.Context, id int) error {
	tx := app.Db.WithContext(ctx).Where("id = ?", id).Select("id,username,phone,nickname,sex,avatar,email,last_login_time,age").Find(&u)
	return tx.Error
}

func (u *User) Update(ctx *gin.Context) error {
	tx := app.Db.WithContext(ctx).Model(&u).Where("id = ?", u.ID).Updates(u)
	return tx.Error
}
