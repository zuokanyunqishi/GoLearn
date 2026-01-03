package user

import (
	baseCommon "GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util"
)

type User struct {
	baseCommon.BaseUsers
	conn   interfaces.Connection // 当前持有的链接
	token  string                // 登陆成功后返回的token
	client *util.Client
}

func (u *User) GetConn() interfaces.Connection {
	return u.conn
}

func (u *User) GetToken() string {
	return u.token
}

func NewUser(client *util.Client) *User {
	return &User{client: client}
}
