package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	serverChatRoom "GoLearn/chat/server/model/chatRoom"
	"GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"net/http"
)

type ChatRoomCreate struct {
	util.BasRouter
}

func (c *ChatRoomCreate) Handle(request interfaces.Request) {
	var req chatRoom.CreateRequest
	c.UnmarshalRequest(request, &req)
	if c.GetHandleErr() != nil {
		util.Response(request.GetConn(),
			chatRoom.CreateResponse{Code: http.StatusBadRequest, ErrMsg: c.GetHandleErr().Error()},
			messageType.ChatRoomCreateResponse)
		return
	}

	u := new(user.User)
	baseUser, err := u.GetUserByToken(req.Token)
	if err != nil {
		util.Response(request.GetConn(),
			chatRoom.CreateResponse{Code: http.StatusBadRequest, ErrMsg: c.GetHandleErr().Error()},
			messageType.ChatRoomCreateResponse)
	}
	s := new(serverChatRoom.ChatRoom)

	req.Room.CreateUser = baseUser
	createResult := s.Create(req.Room)
	var repCode = uint32(http.StatusOK)
	var errMsg string
	if createResult != nil {
		repCode = http.StatusBadRequest
		errMsg = createResult.Error()
	}

	util.Response(request.GetConn(), chatRoom.CreateResponse{
		Code:   repCode,
		ErrMsg: errMsg,
	}, messageType.ChatRoomCreateResponse)

}
