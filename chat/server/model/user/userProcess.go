package user

import (
	"GoLearn/chat/interfaces"
)

type Process struct {
	User *User
	Conn interfaces.Connection
}

func NewUserProcess(user *User, conn interfaces.Connection) *Process {
	return &Process{User: user, Conn: conn}
}
