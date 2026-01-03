package user

import (
	baseCommon "GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/message/login"
	"GoLearn/chat/server/common"
	"GoLearn/chat/util"
	"GoLearn/chat/util/errors"
	"GoLearn/chat/util/zlog"
	"encoding/json"
	"github.com/go-basic/uuid"
	"github.com/syyongx/php2go"
	"sync"
)

// 用户结构体
type User struct {
	UserName string `json:"user_name"`
	UserPwd  string `json:"user_pwd"`
	UserId   uint32 `json:"user_id"`
	NickName string `json:"nick_name"`
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetUserId() uint32 {
	return u.UserId
}

// 获取好友
func (u *User) GetFriendsList(userId uint32) ([]baseCommon.BaseUsers, interfaces.Error) {
	userFriends := make([]baseCommon.BaseUsers, 0)

	userFriendsIds, _ := u.GetFriendsIds(userId)

	for _, id := range userFriendsIds {
		// 判断是否在线
		if userProcess, errTmp := OnlineUser.FindOne(id); errTmp == nil {
			userFriends = append(userFriends, baseCommon.BaseUsers{
				UserName: userProcess.User.UserName,
				UserId:   userProcess.User.UserId,
				Status:   baseCommon.UserOnLine,
				NickName: userProcess.User.NickName,
			})
			continue
		}
		if uFriend, tmpErr := u.GetUserByUserId(id); tmpErr == nil {
			userFriends = append(userFriends, baseCommon.BaseUsers{
				UserName: uFriend.UserName,
				UserId:   uFriend.UserId,
				Status:   baseCommon.UserOffLine,
				NickName: uFriend.NickName,
			})
		}

	}
	return userFriends, nil
}

// 保存用户
func (u *User) Save(request login.RegisterRequest) interfaces.Error {

	u.UserName = request.UserName
	u.UserPwd = php2go.Md5(request.UserPwd)
	u.makeUserId()
	u.NickName = ""

	userStr, err := json.Marshal(u)
	if err != nil {
		return errors.NewNoCode("用户保存失败")
	}

	result := util.Redis.HSet(
		util.Redis.Context(),
		common.ServerUserMapKey,
		u.UserName,
		string(userStr),
	)

	if result.Err() != nil || result.Val() < 1 {
		zlog.Error(result.Err().Error())
		return errors.NewNoCode("用户保存失败")
	}

	util.Redis.HSet(util.Redis.Context(),
		common.UserIdUserNameMapKey,
		util.UnitToString(u.GetUserId()), u.UserName)

	return nil
}

// 根据用户名查询否存在
func (u *User) ExistsByUserName(userName string) bool {
	return util.Redis.HExists(util.Redis.Context(), common.ServerUserMapKey, userName).Val()
}

// 根据用户id 查询是否存在
func (u *User) ExistsByUserId(userId uint32) bool {
	return util.Redis.HExists(util.Redis.Context(), common.UserIdUserNameMapKey, util.UnitToString(userId)).Val()

}

// 自增用户id
func (u *User) makeUserId() uint32 {
	once := new(sync.Once)
	once.Do(func() {
		u.UserId = uint32(util.Redis.Incr(util.Redis.Context(), common.UserIdIncrementKey).Val())
	})

	return u.UserId
}

// 根据用户名获取用户
func (u *User) GetUserByUserName(UserName string) (*User, interfaces.Error) {

	bytes, err := util.Redis.HGet(util.Redis.Context(), common.ServerUserMapKey, UserName).Bytes()
	if err != nil {
		return nil, errors.New("user not found", common.UserNotFound)
	}
	err = json.Unmarshal(bytes, u)
	if err != nil {
		zlog.Error(err.Error())
		return nil, errors.NewNoCode("GetUserByUserName ,  Unmarshal error")
	}
	return u, nil
}

// 根据用户id 获取用户
func (u *User) GetUserByUserId(userId uint32) (*User, interfaces.Error) {
	userName, err := u.GetUserNameByUserId(userId)
	if err != nil {
		return nil, err
	}

	return u.GetUserByUserName(userName)
}

// 根据用户id获取用户实例
func (u *User) GetUserNameByUserId(userId uint32) (string, interfaces.Error) {
	result := util.Redis.HGet(util.Redis.Context(), common.UserIdUserNameMapKey, util.UnitToString(userId))
	err := result.Err()
	if err != nil {
		zlog.Infof("GetUserNameByUserId %s", err.Error())
		return "", errors.New("not found user", common.UserNotFound)
	}
	return result.Val(), nil
}

func (u *User) MakeLoginToken() string {
	return php2go.Md5(php2go.Uniqid("_login_") + uuid.New())
}

func (u *User) SaveToken(token string, userId uint32) {
	util.Redis.HSet(util.Redis.Context(), common.ServerUserTokenMapKey, token, userId)
}

func (u *User) GetUserIdToken(token string) (uint32, interfaces.Error) {
	result := util.Redis.HGet(util.Redis.Context(), common.ServerUserTokenMapKey, token)
	if result.Err() != nil {
		return 0, errors.New("user not found", common.UserNotFound)
	}

	return util.StringToUint32(result.Val()), nil
}

func (u *User) GetFriendsIds(userId uint32) ([]uint32, interfaces.Error) {
	result := util.Redis.HGet(util.Redis.Context(), common.UserFriendsHashMapKey, util.UnitToString(userId))

	var userFriendsIds []uint32
	if result.Err() != nil {
		return userFriendsIds, nil
	}
	bytes, err := result.Bytes()
	if err != nil {
		zlog.Errorf("GetFriendsList result.Bytes error %s", err.Error())
		return userFriendsIds, nil
	}

	_ = json.Unmarshal(bytes, &userFriendsIds)

	return userFriendsIds, nil
}

func (u *User) AddFriends(myUserId uint32, friendsUserIds []uint32) interfaces.Error {
	ids, _ := u.GetFriendsIds(myUserId)
	set := map[uint32]int{}
	for _, id := range ids {
		set[id] = 1
	}
	for _, friendsUserId := range friendsUserIds {
		if u.ExistsByUserId(friendsUserId) {
			set[friendsUserId] = 1
		}
	}

	var userFriendIds []uint32
	for id, _ := range set {
		userFriendIds = append(userFriendIds, id)
	}

	if len(userFriendIds) > len(ids) {
		bytes, _ := json.Marshal(userFriendIds)

		result := util.Redis.HSet(util.Redis.Context(), common.UserFriendsHashMapKey, myUserId, string(bytes))

		if result.Err() != nil || result.Val() < 1 {
			zlog.Warnf("%d 用户添加好友 错误 %s", myUserId, result.Err())
			return errors.NewNoCode("好友添加失败。。")
		}
	}

	return nil

}

func (u *User) GetUserByToken(token string) (baseCommon.BaseUsers, interfaces.Error) {

	var baseUser baseCommon.BaseUsers
	userId, err := u.GetUserIdToken(token)
	if err != nil {
		return baseUser, err
	}
	u, err = u.GetUserByUserId(userId)

	if err != nil {
		return baseUser, err
	}

	baseUser.NickName = u.NickName
	baseUser.UserId = u.UserId
	baseUser.Status = baseCommon.UserOnLine
	return baseUser, nil
}
