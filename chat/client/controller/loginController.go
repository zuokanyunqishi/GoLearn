package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/util"
)

type LoginResponseRoute struct {
	util.BasRouter
	response login.ResponseLogin
}

// 处理业务中前钩子
func (r *LoginResponseRoute) Handle(request interfaces.Request) {
	var response login.ResponseLogin
	r.UnmarshalRequest(request, &response)
	if r.GetHandleErr() != nil {
		return
	}

	r.response = response
	r.SetHandleOK(true)
	r.SendHandleOk()

}

// 处理业务后钩子
func (r *LoginResponseRoute) PostHandle(request interfaces.Request) {
}

func (r *LoginResponseRoute) GetResponse() (login.ResponseLogin, interfaces.Error) {

	err := r.HandleResponse()
	if err != nil {
		return login.ResponseLogin{}, err
	}

	return r.response, nil
}

func (r *LoginResponseRoute) AfterBusinessHandle() {
	r.BasRouter.AfterBusinessHandle()
	r.response = login.ResponseLogin{}
}
