package chatRoom

import (
	"GoLearn/chat/common"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/server/model/user"
	"encoding/json"
	"net/http"
	"sync"
)

type ServerRoom struct {
	common.ChatRoom
	RoomUserProcess map[uint32]*user.Process
	lock            sync.Mutex
}

func NewServerRoom(chatRoom common.ChatRoom) *ServerRoom {
	return &ServerRoom{ChatRoom: chatRoom,
		RoomUserProcess: make(map[uint32]*user.Process)}
}

func (sr *ServerRoom) AddRoomUserProcess(process *user.Process) {
	sr.lock.Lock()
	defer sr.lock.Unlock()
	sr.RoomUserProcess[process.User.UserId] = process
}

func (sr *ServerRoom) RemoveUser(userId uint32) {
	sr.lock.Lock()
	defer sr.lock.Unlock()
	if _, ok := sr.RoomUserProcess[userId]; ok {
		delete(sr.RoomUserProcess, userId)
	}
}

func (sr *ServerRoom) SendMessage(sendUserId uint32, message string) {
	sendUser, _ := (new(user.User)).GetUserByUserId(sendUserId)
	response := chatRoom.RoomMessageReceive{
		Code:    http.StatusOK,
		Message: message,
		User: common.BaseUsers{
			UserName: sendUser.UserName,
			UserId:   sendUser.UserId,
			Status:   1,
			NickName: sendUser.NickName,
		},
		ErrMsg: "",
	}
	b, _ := json.Marshal(response)
	for _, v := range sr.RoomUserProcess {
		v.Conn.SendMsg(messageType.ChatRoomMessageReceive, b)
	}
}
