package user

import (
	"GoLearn/chat/client/route"
	baseCommon "GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/message/online"
	"encoding/json"
)

func (u *User) GetOnlineUsers() ([]baseCommon.BaseUsers, interfaces.Error) {
	onLineUsers := make([]baseCommon.BaseUsers, 0)
	// 渲染当前在线的人
	req := online.UserListRequest{Token: u.GetToken()}
	reqBytes, _ := json.Marshal(req)
	u.GetConn().SendMsg(messageType.OnlineUserListRequest, reqBytes)
	response, err := route.OnlineUsersController.GetResponse()

	if err != nil {
		return onLineUsers, err
	}

	return response.UserList, nil

}
