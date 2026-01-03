package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	serverRoom "GoLearn/chat/server/model/chatRoom"
	serverUser "GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"fmt"
)

type IntoRoomRouter struct {
	util.BasRouter
}

func (ir *IntoRoomRouter) Handle(request interfaces.Request) {

	var req chatRoom.IntoRoomRequest
	ir.UnmarshalRequest(request, &req)
	if ir.GetHandleErr() != nil {
		fmt.Println(ir.GetHandleErr().Error())
		return
	}

	u := new(serverUser.User)
	baseUser, _ := u.GetUserByToken(req.Token)
	room, _ := serverRoom.FindOne(req.RoomId)

	var mgrRoom *serverRoom.ServerRoom
	var err interfaces.Error
	mgrRoom, err = serverRoom.ServerRoomMgr.FindOne(room.RoomId)
	userProcess, _ := serverUser.OnlineUser.FindOne(baseUser.UserId)

	if err != nil {
		mgrRoom = serverRoom.NewServerRoom(room)
		serverRoom.ServerRoomMgr.Add(mgrRoom)
	}

	mgrRoom.AddRoomUserProcess(userProcess)

	util.Response(request.GetConn(), chatRoom.IntoRoomResponse{
		Code: 200,
		Room: room,
	}, messageType.IntoRoomResponse)
}
