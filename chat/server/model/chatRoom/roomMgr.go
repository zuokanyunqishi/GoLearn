package chatRoom

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/errors"
	"sync"
)

//--------------------------------------------------------------//

var ServerRoomMgr = NewChatRoomsMgr()

type ChatRoomsMgr struct {
	Rooms map[uint32]*ServerRoom
	lock  sync.RWMutex
}

func (cm *ChatRoomsMgr) Add(room *ServerRoom) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.Rooms[room.RoomId] = room
}

func (cm *ChatRoomsMgr) Remove(roomId uint32) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	delete(cm.Rooms, roomId)
}

func (cm *ChatRoomsMgr) FindOne(roomId uint32) (*ServerRoom, interfaces.Error) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	room, ok := cm.Rooms[roomId]
	if !ok {
		return nil, errors.NewNoCode("not found room")
	}

	return room, nil
}

func (cm *ChatRoomsMgr) SendRoomMessage(roomId, sendUserId uint32, message string) {
	room, err := cm.FindOne(roomId)
	if err != nil {
		return
	}
	room.SendMessage(sendUserId, message)
}

func NewChatRoomsMgr() *ChatRoomsMgr {
	return &ChatRoomsMgr{Rooms: make(map[uint32]*ServerRoom)}
}
