package user

import (
	clientErr "GoLearn/chat/client/errors"
	"GoLearn/chat/client/route"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/message/messageType"
	errorsNew "GoLearn/chat/util/errors"
	"encoding/json"
	"net/http"
)

func (u *User) Login(userName, userPwd string) (error interfaces.Error) {

	requestBytes, err := json.Marshal(login.RequestLogin{
		UserName: userName,
		UserPwd:  userPwd,
	})

	if err != nil {
		return clientErr.ErrLoginRequestMarshal
	}
	conn := u.client.CreateConnection(false)

	conn.SendMsg(messageType.LoginRequest, requestBytes)

	rc := route.LoginResponseController
	// 重置login response
	defer rc.AfterBusinessHandle()
	response, responseErr := rc.GetResponse()

	defer func() {
		if error != nil {
			u.client.GetConnManager().RemoveConn(conn)
			conn.Stop()
		}
	}()

	if responseErr != nil {
		error = responseErr
		return
	}

	if response.Code != http.StatusOK {
		error = errorsNew.New(response.ErrMsg, clientErr.ErrCodeLoginResponse)
		return
	}
	u.token = response.Token
	u.conn = conn
	u.BaseUsers = response.User

	return nil
}
