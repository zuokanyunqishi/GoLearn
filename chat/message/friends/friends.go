package friends

import (
	"GoLearn/chat/common"
)

// 用户好友响应
type UserFriendsResponse struct {
	Code            uint32 `json:"code"`
	ErrMsg          string `json:"errMsg"`
	UserFriendsList []common.BaseUsers
}

// 用户好友列表请求
type UserFriendsRequest struct {
	Token string `json:"token"`
	Page  uint   `json:"page"`
}

// 添加好友请求
type FriendsAdd struct {
	Token          string   `json:"token"`
	FriendsUserIds []uint32 `json:"friends_user_ids"`
}

// 添加好友请求
type FriendsAddResponse struct {
	Code   uint32 `json:"code"`
	ErrMsg string `json:"errMsg"`
	Data   string `json:"data"`
}
