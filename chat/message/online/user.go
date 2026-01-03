package online

import baseCommon "GoLearn/chat/common"

// 在线用户列表请求体
type UserListRequest struct {
	Token string `json:"token"`
}

// 在线用户列表响应
type UserListResponse struct {
	Code     uint32 `json:"code"`
	ErrMsg   string `json:"errMsg"`
	UserList []baseCommon.BaseUsers
}
