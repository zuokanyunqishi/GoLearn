package user

import (
	"GoLearn/chat/client/route"
	"GoLearn/chat/common"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	"encoding/json"
)

func (u *User) SendMessage(roomId uint32, message string) bool {

	roomMessage := chatRoom.RoomMessageSend{
		ToRoomId:   roomId,
		Message:    message,
		Token:      u.GetToken(),
		SendUserId: u.UserId,
	}
	b, _ := json.Marshal(roomMessage)
	_ = u.GetConn().SendMsg(messageType.ChatRoomMessageSend, b)
	return true
}

func (u *User) IntoRoom(roomId uint32) (common.ChatRoom, error) {
	req := chatRoom.IntoRoomRequest{
		Token:  u.GetToken(),
		RoomId: roomId,
	}

	b, _ := json.Marshal(req)
	_ = u.GetConn().SendMsg(messageType.IntoRoomRequest, b)
	response, err := route.IntoRoomController.GetResponse()

	if err != nil {
		return common.ChatRoom{}, err
	}

	return response.Room, nil
}

func (u *User) ReceiveRoomMessages() chan chatRoom.RoomMessageReceive {

	return route.RoomMessageReceiveController.GetRoomMessage()
}
