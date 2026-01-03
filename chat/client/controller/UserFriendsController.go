package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/friends"
	"GoLearn/chat/util"
	"GoLearn/chat/util/zlog"
)

type UserFriendsRoute struct {
	util.BasRouter
	response friends.UserFriendsResponse
}

// 处理业务中前钩子
func (f *UserFriendsRoute) Handle(request interfaces.Request) {
	var response friends.UserFriendsResponse
	f.UnmarshalRequest(request, &response)
	if f.GetHandleErr() != nil {
		zlog.Error("user friends response err" + f.GetHandleErr().Error())
		return
	}

	f.response = response
	f.SetHandleOK(true)
	f.SendHandleOk()

}

func (f *UserFriendsRoute) GetResponse() (friends.UserFriendsResponse, interfaces.Error) {

	err := f.HandleResponse()

	if err != nil {
		return friends.UserFriendsResponse{}, err
	}
	return f.response, nil
}

func (f *UserFriendsRoute) AfterBusinessHandle() {
	f.BasRouter.AfterBusinessHandle()
	f.response = friends.UserFriendsResponse{}
}

// --------------------------------------------------------------------------------------
type FriendsAddRoute struct {
	util.BasRouter
	response friends.FriendsAddResponse
}

func (f *FriendsAddRoute) Handle(request interfaces.Request) {
	var response friends.FriendsAddResponse
	f.UnmarshalRequest(request, &response)
	if f.GetHandleErr() != nil {
		zlog.Error("user friends response err" + f.GetHandleErr().Error())
		return
	}
	f.response = response
	f.SetHandleOK(true)
	f.SendHandleOk()

}

func (f *FriendsAddRoute) AfterBusinessHandle() {
	f.BasRouter.AfterBusinessHandle()
	f.response = friends.FriendsAddResponse{}
}

func (f *FriendsAddRoute) GetResponse() (friends.FriendsAddResponse, interfaces.Error) {

	err := f.HandleResponse()

	if err != nil {
		return friends.FriendsAddResponse{}, err
	}
	return f.response, nil
}
