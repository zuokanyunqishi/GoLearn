package user

import (
	clientErr "GoLearn/chat/client/errors"
	"GoLearn/chat/client/route"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/util/errors"
	"GoLearn/chat/util/zlog"
	"encoding/json"
)

func (u *User) Register(userName, userPwd, ConfirmPwd string) (error interfaces.Error) {

	conn := u.client.CreateConnection(false)
	var response login.RegisterResponse
	regReuestBytes, _ := json.Marshal(login.RegisterRequest{
		UserName:   userName,
		UserPwd:    userPwd,
		ConfirmPwd: ConfirmPwd,
	})

	err := conn.SendMsg(messageType.LoginRegisterRequest, regReuestBytes)

	if err != nil {
		zlog.Error("Register send msg error" + err.Error())
		return error
	}
	response, error = route.RegisterResponseController.GetResponse()
	defer func() {
		u.client.GetConnManager().RemoveConn(conn)
		conn.Stop()
	}()
	if error != nil {
		zlog.PrintError(error.Error())
		return
	}

	if response.Code != 200 {
		return errors.New(response.ErrMsg, clientErr.ErrCodeRegisterResponse)
	}
	return nil
}
