package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/chatRoom"
	"GoLearn/chat/util"
)

type ChatRoomListRoute struct {
	util.BasRouter
	response chatRoom.ChatRoomListResponse
}

func (c *ChatRoomListRoute) Handle(request interfaces.Request) {
	var rep chatRoom.ChatRoomListResponse
	c.UnmarshalRequest(request, &rep)
	if c.GetHandleErr() != nil {
		return
	}

	// 处理响应
	c.response = rep
	c.SetHandleOK(true)
	c.SendHandleOk()
}

func (c *ChatRoomListRoute) GetResponse() (chatRoom.ChatRoomListResponse, interfaces.Error) {
	err := c.HandleResponse()
	if err != nil {
		return chatRoom.ChatRoomListResponse{}, err
	}

	return c.response, nil
}

func (c *ChatRoomListRoute) AfterBusinessHandle() {
	c.BasRouter.AfterBusinessHandle()
	c.response = chatRoom.ChatRoomListResponse{}
}

// 																								//
/*-----------------------------------------------------------------------------------------------*/
//																								//

// 创建聊天室
type ChatRoomCreateRoute struct {
	util.BasRouter
	response chatRoom.CreateResponse
}

func (c *ChatRoomCreateRoute) Handle(request interfaces.Request) {
	var rep chatRoom.CreateResponse
	c.UnmarshalRequest(request, &rep)
	if c.GetHandleErr() != nil {
		return
	}

	// 处理响应
	c.response = rep
	c.SetHandleOK(true)
	c.SendHandleOk()
}

func (c *ChatRoomCreateRoute) GetResponse() (chatRoom.CreateResponse, interfaces.Error) {
	err := c.HandleResponse()
	if err != nil {
		return chatRoom.CreateResponse{}, err
	}

	return c.response, nil
}

func (c *ChatRoomCreateRoute) AfterBusinessHandle() {
	c.BasRouter.AfterBusinessHandle()
	c.response = chatRoom.CreateResponse{}
}

//------------------------------------------------------------------------------------------//

// 进入聊天室路由
type IntoRoomRoute struct {
	util.BasRouter
	response chatRoom.IntoRoomResponse
}

func (I *IntoRoomRoute) Handle(request interfaces.Request) {
	var rep chatRoom.IntoRoomResponse
	I.UnmarshalRequest(request, &rep)
	if I.GetHandleErr() != nil {
		return
	}

	// 处理响应
	I.response = rep
	I.SetHandleOK(true)
	I.SendHandleOk()
}

func (I *IntoRoomRoute) GetResponse() (chatRoom.IntoRoomResponse, interfaces.Error) {
	err := I.HandleResponse()
	if err != nil {
		return chatRoom.IntoRoomResponse{}, err
	}
	return I.response, nil
}

func (I *IntoRoomRoute) AfterBusinessHandle() {
	I.BasRouter.AfterBusinessHandle()
	I.response = chatRoom.IntoRoomResponse{}
}

type RoomMessageReceiveRoute struct {
	util.BasRouter
	response    chatRoom.RoomMessageReceive
	messageChan chan chatRoom.RoomMessageReceive
}

func (m *RoomMessageReceiveRoute) Handle(request interfaces.Request) {
	if m.messageChan == nil {
		m.messageChan = make(chan chatRoom.RoomMessageReceive)
	}

	var rep chatRoom.RoomMessageReceive
	m.UnmarshalRequest(request, &rep)
	if m.GetHandleErr() != nil {
		return
	}
	m.messageChan <- rep
}

func (m *RoomMessageReceiveRoute) GetRoomMessage() chan chatRoom.RoomMessageReceive {
	return m.messageChan
}
