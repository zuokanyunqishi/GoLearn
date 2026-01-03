package exceptions

import "gopkg.in/errgo.v2/errors"

var (
	BIG_ERROR = errors.New("test")
	ERROR2    = errors.New("test")
	ERROR3    = errors.New("test")

	ErrUserNotFound  = errors.New("用户找不到")
	ErrWrongPassword = errors.New("密码错误")
	JwtTokenExpire   = errors.New("jwt token expired")
	JwtTokenInvalid  = errors.New("jwt token is invalid")
	UserNameIsUsed   = errors.New("用户名已经占用")
)
