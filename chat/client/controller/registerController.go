package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/util"
	"GoLearn/chat/util/zlog"
)

type RegisterResponseRoute struct {
	util.BasRouter
	response login.RegisterResponse
}

// 处理业务中前钩子
func (r *RegisterResponseRoute) Handle(request interfaces.Request) {
	var response login.RegisterResponse
	r.UnmarshalRequest(request, &response)

	if r.GetHandleErr() != nil {
		zlog.Error("register err " + r.GetHandleErr().Error())
		return
	}

	r.response = response
	r.SetHandleOK(true)
	r.SendHandleOk()
}

func (r *RegisterResponseRoute) GetResponse() (login.RegisterResponse, interfaces.Error) {

	err := r.HandleResponse()

	if err != nil {
		return login.RegisterResponse{}, err
	}
	return r.response, nil
}

func (r *RegisterResponseRoute) AfterBusinessHandle() {
	r.BasRouter.AfterBusinessHandle()
	r.response = login.RegisterResponse{}
}
