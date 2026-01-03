package user

import (
	"GoLearn/chat/client/route"
	"GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/friends"
	"GoLearn/chat/message/messageType"
	"GoLearn/chat/util/errors"
	"encoding/json"
	"net/http"
)

func (u *User) GetFriends() (userFriendsList []common.BaseUsers, err interfaces.Error) {

	var request friends.UserFriendsRequest
	request.Page = 0
	request.Token = u.token
	data, errTmp := json.Marshal(request)
	if errTmp != nil {
		return userFriendsList, errors.NewNoCode(errTmp.Error())
	}
	u.GetConn().SendMsg(messageType.UserFriendsListRequest, data)
	response, errTmp := route.UserFriendsController.GetResponse()
	defer route.UserFriendsController.AfterBusinessHandle()
	if errTmp != nil {
		return userFriendsList, errors.NewNoCode(errTmp.Error())
	}

	return response.UserFriendsList, nil

}

// 添加好友
func (u *User) AddFriend(userId ...uint32) interfaces.Error {
	var request friends.FriendsAdd
	request.Token = u.GetToken()
	request.FriendsUserIds = userId
	bytes, tmpErr := json.Marshal(request)
	if tmpErr != nil {
		return errors.NewNoCode(tmpErr.Error())
	}
	u.GetConn().SendMsg(messageType.FriendsAddRequest, bytes)
	response, err := route.FriendsAddController.GetResponse()
	if err != nil {
		return errors.NewNoCode(err.Error())
	}

	if response.Code != http.StatusOK {
		return errors.NewNoCode(response.ErrMsg)
	}
	return nil
}
