package common

type BaseUsers struct {
	UserName string `json:"user_name"`
	UserId   uint32 `json:"user_id"`   // 用户id
	Status   uint8  `json:"status"`    // 状态
	NickName string `json:"nick_name"` // 昵称
}

// ChatRoom 聊天室结构体
type ChatRoom struct {
	RoomId       uint32            `json:"room_id"`       // 聊天室id
	RoomName     string            `json:"room_name"`     // 聊天室名字
	RoomDescribe string            `json:"room_describe"` // 聊天室描述
	CreateUser   BaseUsers         `json:"create_user"`   // 聊天室创建者
	RoomUsers    map[uint32]string `json:"room_users"`    // 聊天室成员 key为用户id
}
