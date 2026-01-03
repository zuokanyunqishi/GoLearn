package controller

import (
	"GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"encoding/json"
	"github.com/syyongx/php2go"
	"net/http"
)

// 登陆
type Login struct {
	util.BasRouter
}

func (b *Login) Handle(request interfaces.Request) {

	var loginRequest login.RequestLogin
	var err error

	err = json.Unmarshal(request.GetData(), &loginRequest)

	if err != nil {
		util.Response(request.GetConn(), login.ResponseLogin{
			Code:   http.StatusUnprocessableEntity,
			ErrMsg: "数据错误",
			Token:  "",
		}, messageType.LoginResponse)
		return
	}

	u := new(user.User)
	_, err = u.GetUserByUserName(loginRequest.UserName)
	if err != nil {
		util.Response(request.GetConn(), login.ResponseLogin{
			Code:   http.StatusForbidden,
			ErrMsg: "userName or userPwd err",
			Token:  "",
		}, messageType.LoginResponse)
		return
	}
	if u.UserPwd != php2go.Md5(loginRequest.UserPwd) {
		util.Response(request.GetConn(), login.ResponseLogin{
			Code:   http.StatusForbidden,
			ErrMsg: "userName or userPwd err",
			Token:  "",
		}, messageType.LoginResponse)
		return
	}
	token := u.MakeLoginToken()
	u.SaveToken(token, u.GetUserId())
	user.OnlineUser.Add(
		user.NewUserProcess(u, request.GetConn()),
	)

	util.Response(request.GetConn(), login.ResponseLogin{
		Code:   http.StatusOK,
		ErrMsg: "成功",
		Token:  token,
		User: common.BaseUsers{
			UserName: u.UserName,
			UserId:   u.UserId,
			Status:   common.UserOnLine,
			NickName: u.NickName,
		},
	}, messageType.LoginResponse)
}
