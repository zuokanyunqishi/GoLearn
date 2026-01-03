package chatRoom

import "GoLearn/chat/common"

// 聊天室请求
type ChatRoomListRequest struct {
	Token string `json:"token"`
}

// 聊天室列表响应
type ChatRoomListResponse struct {
	Code   uint32            `json:"code"`
	ErrMsg string            `json:"errMsg"`
	List   []common.ChatRoom `json:"list"`
}

// 创建聊天室请求
type CreateRequest struct {
	Room  common.ChatRoom `json:"room"`
	Token string          `json:"token"`
}

// 创建聊天室响应
type CreateResponse struct {
	Code   uint32 `json:"code"`
	ErrMsg string `json:"errMsg"`
}

// 聊天室发送message
type RoomMessageSend struct {
	ToRoomId   uint32 `json:"to_room_id"`
	SendUserId uint32 `json:"send_user_id"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}

// 聊天室响应message
type RoomMessageReceive struct {
	Code    uint32           `json:"code"`
	Message string           `json:"message"`
	User    common.BaseUsers `json:"user"`
	ErrMsg  string           `json:"err_msg"`
}

// 进入聊天室请求
type IntoRoomRequest struct {
	Token  string `json:"token"`
	RoomId uint32 `json:"room_id"`
}

// 进入聊天室响应
type IntoRoomResponse struct {
	Code uint32          `json:"code"`
	Room common.ChatRoom `json:"room"`
}
