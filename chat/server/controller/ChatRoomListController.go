package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	serverChatRoom "GoLearn/chat/server/model/chatRoom"
	"GoLearn/chat/util"
	"net/http"
)

type ChatRoomListRoute struct {
	util.BasRouter
}

func (c *ChatRoomListRoute) Handle(request interfaces.Request) {
	var req chatRoom.ChatRoomListResponse
	c.UnmarshalRequest(request, &req)
	if c.GetHandleErr() != nil {
		util.Response(request.GetConn(), chatRoom.ChatRoomListResponse{
			Code:   http.StatusBadRequest,
			ErrMsg: c.GetHandleErr().Error(),
			List:   nil,
		}, messageType.ChatRoomListResponse)
		return
	}
	list, err := new(serverChatRoom.ChatRoom).List()
	if err != nil {
		util.Response(request.GetConn(),
			chatRoom.ChatRoomListResponse{
				Code:   http.StatusBadRequest,
				ErrMsg: err.Error(),
				List:   nil,
			}, messageType.ChatRoomListResponse,
		)
		return
	}

	util.Response(request.GetConn(),
		chatRoom.ChatRoomListResponse{
			Code:   http.StatusOK,
			ErrMsg: "",
			List:   list,
		}, messageType.ChatRoomListResponse)

}
