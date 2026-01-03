package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/message/online"
	"GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"net/http"
)

type OnlineUserRoute struct {
	util.BasRouter
}

func (u *OnlineUserRoute) Handle(request interfaces.Request) {

	var onlineRq online.UserListRequest

	u.UnmarshalRequest(request, &onlineRq)

	if u.GetHandleErr() != nil {
		util.Response(request.GetConn(), online.UserListResponse{
			Code:     http.StatusBadRequest,
			ErrMsg:   u.GetHandleErr().Error(),
			UserList: nil,
		}, messageType.OnlineUserListResponse)
		return
	}

	userModel := &user.User{}

	// 校验token
	_, err := userModel.GetUserIdToken(onlineRq.Token)
	if err != nil {
		util.Response(
			request.GetConn(),
			online.UserListResponse{Code: http.StatusForbidden, ErrMsg: err.Error()},
			messageType.OnlineUserListResponse,
		)
		return
	}

	util.Response(
		request.GetConn(),
		online.UserListResponse{Code: http.StatusOK, UserList: user.OnlineUser.GetAll()},
		messageType.OnlineUserListResponse,
	)

}
