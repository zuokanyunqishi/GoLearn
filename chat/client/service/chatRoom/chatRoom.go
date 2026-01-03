package chatRoom

import (
	"GoLearn/chat/client/route"
	"GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/util/errors"
	"encoding/json"
	"net/http"
)

// 聊天室列表
func ChatRooms(conn interfaces.Connection, token string) ([]common.ChatRoom, interfaces.Error) {
	var req chatRoom.ChatRoomListRequest
	req.Token = token
	marshal, _ := json.Marshal(req)
	conn.SendMsg(messageType.ChatRoomListRequest, marshal)
	response, err := route.ChatRoomListController.GetResponse()
	if err != nil {
		return nil, errors.NewNoCode(err.Error())
	}

	if response.Code != http.StatusOK {
		return nil, errors.NewNoCode(response.ErrMsg)
	}

	return response.List, nil
}

func CreateRoom(conn interfaces.Connection, token string, room common.ChatRoom) interfaces.Error {
	var req chatRoom.CreateRequest
	req.Token = token
	req.Room = room
	bytes, _ := json.Marshal(req)
	conn.SendMsg(messageType.ChatRoomCreateRequest, bytes)

	response, err := route.ChatRoomCreateController.GetResponse()
	if err != nil {
		return errors.NewNoCode(err.Error())
	}

	if response.Code != http.StatusOK {
		return errors.NewNoCode(response.ErrMsg)
	}

	return nil
}
