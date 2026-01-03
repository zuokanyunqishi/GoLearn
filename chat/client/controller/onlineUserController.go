package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/online"
	"GoLearn/chat/util"
	"GoLearn/chat/util/zlog"
)

type OnlineUsersRoute struct {
	util.BasRouter
	response online.UserListResponse
}

func (f *OnlineUsersRoute) Handle(request interfaces.Request) {
	var response online.UserListResponse
	f.UnmarshalRequest(request, &response)
	if f.GetHandleErr() != nil {
		zlog.Error("user friends response err" + f.GetHandleErr().Error())
		return
	}
	f.response = response
	f.SetHandleOK(true)
	f.SendHandleOk()

}

func (f *OnlineUsersRoute) AfterBusinessHandle() {
	f.BasRouter.AfterBusinessHandle()
	f.response = online.UserListResponse{}
}

func (f *OnlineUsersRoute) GetResponse() (online.UserListResponse, interfaces.Error) {

	err := f.HandleResponse()

	if err != nil {
		return online.UserListResponse{}, err
	}
	return f.response, nil
}
