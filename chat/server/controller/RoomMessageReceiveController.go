package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	serverRoom "GoLearn/chat/server/model/chatRoom"
	"GoLearn/chat/util"
	"net/http"
)

type RoomMessageReceive struct {
	util.BasRouter
}

func (r *RoomMessageReceive) Handle(request interfaces.Request) {
	var req chatRoom.RoomMessageSend
	r.UnmarshalRequest(request, &req)
	if r.GetHandleErr() != nil {
		util.Response(request.GetConn(), chatRoom.RoomMessageReceive{
			Code:   http.StatusBadRequest,
			ErrMsg: r.GetHandleErr().Error(),
		}, messageType.ChatRoomMessageReceive)
	}
	serverRoom.ServerRoomMgr.SendRoomMessage(req.ToRoomId, req.SendUserId, req.Message)
}
