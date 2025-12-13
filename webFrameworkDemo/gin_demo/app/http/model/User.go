package model

import (
	"context"
	app "speed/bootstrap"
)

type User struct {
	ID            int    `gorm:"primaryKey;autoIncrement;column:id"`        // 主键自增[1](@ref)[6](@ref)
	UserName      string `gorm:"type:text(30);not null;column:user_name"`   // 用户名，非空[2](@ref)[6](@ref)
	Phone         string `gorm:"type:text(11);not null;column:phone"`       // 手机号，非空[2](@ref)[7](@ref)
	Password      string `gorm:"type:text(255);not null;column:password"`   // 密码，非空[2](@ref)[7](@ref)
	Token         string `gorm:"type:text(500);default:'';column:token"`    // Token，默认空字符串[6](@ref)[7](@ref)
	LastLoginTime int64  `gorm:"not null;default:0;column:last_login_time"` // 最后登录时间，默认0[6](@ref)[7](@ref)
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) GetUserByUserName(ctx context.Context, username string) error {
	tx := app.Db.WithContext(ctx).Where("user_name = ?", username).Find(&u)
	return tx.Error
}

func (u *User) Add(ctx context.Context) error {
	tx := app.Db.WithContext(ctx).Save(&u)
	return tx.Error
}
