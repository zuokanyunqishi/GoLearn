package exceptions

import "gopkg.in/errgo.v2/errors"

var (
	BIG_ERROR = errors.New("test")
	ERROR2    = errors.New("test")
	ERROR3    = errors.New("test")

	ErrUserNotFound  = errors.New("用户找不到")
	ErrWrongPassword = errors.New("密码错误")
)
