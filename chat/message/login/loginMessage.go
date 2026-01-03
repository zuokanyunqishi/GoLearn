package login

import baseCommon "GoLearn/chat/common"

// 登录结构体
type RequestLogin struct {
	UserName string `json:"userName"`
	UserPwd  string `json:"userPwd"`
}

// 登录响应
type ResponseLogin struct {
	Code   uint32 `json:"code"`
	ErrMsg string `json:"errMsg"`
	Token  string `json:"token"`
	User   baseCommon.BaseUsers
}

// 注册请求体
type RegisterRequest struct {
	UserName   string `json:"userName"`
	UserPwd    string `json:"userPwd"`
	ConfirmPwd string `json:"confirmPwd"`
}

// 注册响应
type RegisterResponse struct {
	Code   uint32 `json:"code"`
	ErrMsg string `json:"errMsg"`
	Data   string `json:"data"`
}
