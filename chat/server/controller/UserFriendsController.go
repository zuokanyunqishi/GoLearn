package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/friends"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"net/http"
)

type UserFriendsList struct {
	util.BasRouter
}

func (f *UserFriendsList) Handle(request interfaces.Request) {
	var re friends.UserFriendsRequest
	f.UnmarshalRequest(request, &re)

	userModel := &user.User{}

	// 校验token
	userId, err := userModel.GetUserIdToken(re.Token)
	if err != nil {
		util.Response(
			request.GetConn(),
			friends.UserFriendsResponse{Code: http.StatusForbidden, ErrMsg: err.Error()},
			messageType.UserFriendsListResponse,
		)
		return
	}
	userFriends, err := userModel.GetFriendsList(userId)
	if err != nil {
		util.Response(request.GetConn(),
			friends.UserFriendsResponse{
				Code:   http.StatusNotFound,
				ErrMsg: err.Error(),
			}, messageType.UserFriendsListResponse,
		)
		return
	}

	//
	util.Response(request.GetConn(),
		friends.UserFriendsResponse{
			Code:            http.StatusNotFound,
			ErrMsg:          "",
			UserFriendsList: userFriends,
		}, messageType.UserFriendsListResponse)
}

type FriendsAdd struct {
	util.BasRouter
}

func (a *FriendsAdd) Handle(request interfaces.Request) {

	var re friends.FriendsAdd
	a.UnmarshalRequest(request, &re)
	userModel := &user.User{}

	// 校验token
	userId, err := userModel.GetUserIdToken(re.Token)
	if err != nil {
		util.Response(
			request.GetConn(),
			friends.UserFriendsResponse{Code: http.StatusForbidden, ErrMsg: err.Error()},
			messageType.FriendsAddResponse,
		)
		return
	}
	err = userModel.AddFriends(userId, re.FriendsUserIds)
	if err != nil {
		util.Response(
			request.GetConn(),
			friends.FriendsAddResponse{Code: http.StatusBadRequest, ErrMsg: err.Error()},
			messageType.FriendsAddResponse,
		)
		return
	}

	util.Response(
		request.GetConn(),
		friends.FriendsAddResponse{Code: http.StatusOK},
		messageType.FriendsAddResponse)
}
