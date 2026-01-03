package controller

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"encoding/json"
	"net/http"
)

type Register struct {
	util.BasRouter
}

func (r *Register) Handle(request interfaces.Request) {

	var err error
	var registerRequest login.RegisterRequest

	err = json.Unmarshal(request.GetData(), &registerRequest)
	if err != nil {
		util.Response(request.GetConn(), login.RegisterResponse{
			Code:   http.StatusUnprocessableEntity,
			ErrMsg: "数据错误",
			Data:   "",
		}, messageType.LoginRegisterResponse)
		return
	}
	if registerRequest.UserPwd != registerRequest.ConfirmPwd {
		util.Response(request.GetConn(), login.RegisterResponse{
			Code:   http.StatusForbidden,
			ErrMsg: "密码和确认密码不一致",
			Data:   "",
		}, messageType.LoginRegisterResponse)
		return
	}

	u := &user.User{}
	if u.ExistsByUserName(registerRequest.UserName) {
		util.Response(
			request.GetConn(),
			login.RegisterResponse{
				Code:   http.StatusForbidden,
				ErrMsg: "此用户名已经注册,请换一个",
				Data:   "",
			},
			messageType.LoginRegisterResponse,
		)
		return
	}

	if err := u.Save(registerRequest); err != nil {
		util.Response(request.GetConn(), login.RegisterResponse{
			Code:   http.StatusBadRequest,
			ErrMsg: err.Error(),
			Data:   "",
		}, messageType.LoginRegisterResponse)
		return
	}

	util.Response(request.GetConn(), login.RegisterResponse{
		Code:   http.StatusOK,
		ErrMsg: "注册成功",
		Data:   "",
	}, messageType.LoginRegisterResponse)

}
