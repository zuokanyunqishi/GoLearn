package chatRoom

import (
	baseCommon "GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/server/common"
	"GoLearn/chat/server/model/user"
	"GoLearn/chat/util"
	"GoLearn/chat/util/errors"
	"GoLearn/chat/util/zlog"
	"encoding/json"
	"strconv"
	"sync"
)

// 聊天室
type ChatRoom struct {
	RoomId       uint32 `json:"room_id"`        // 聊天室id
	RoomName     string `json:"room_name"`      // 聊天室名字
	RoomDescribe string `json:"room_describe"`  // 聊天室描述
	CreateUserId uint32 `json:"create_user_id"` // 聊天室创建者
}

// 聊天室列表
func (c *ChatRoom) List() ([]baseCommon.ChatRoom, interfaces.Error) {
	var list []baseCommon.ChatRoom
	result := util.Redis.HGetAll(util.Redis.Context(), common.ChatRoomHashMapKey)
	m, err := result.Result()
	if err != nil {
		zlog.Errorf("ChatRoom:list %s", err.Error())
		return nil, errors.NewNoCode("服务错误,请稍后..")
	}
	uModel := new(user.User)

	for _, value := range m {
		var cr baseCommon.ChatRoom
		_ = json.Unmarshal([]byte(value), &cr)

		u, err := uModel.GetUserByUserId(cr.CreateUser.UserId)

		if err != nil {
			continue
		}

		status := baseCommon.UserOnLine
		if !user.OnlineUser.IsOnline(u.UserId) {
			status = baseCommon.UserOffLine
		}

		list = append(list, baseCommon.ChatRoom{
			RoomId:       cr.RoomId,
			RoomName:     cr.RoomName,
			RoomDescribe: cr.RoomDescribe,
			CreateUser: baseCommon.BaseUsers{
				UserName: u.UserName,
				UserId:   u.UserId,
				Status:   uint8(status),
				NickName: "",
			},
		})
	}

	return list, nil

}

func (c *ChatRoom) Create(room baseCommon.ChatRoom) interfaces.Error {
	room.RoomId = c.makeRoomId()
	bytes, _ := json.Marshal(room)
	result := util.Redis.HSet(util.Redis.Context(), common.ChatRoomHashMapKey, room.RoomId, string(bytes))
	if result.Err() != nil || result.Val() < 1 {
		return errors.NewNoCode("创建聊天是失败，请稍后再试。。")
	}

	return nil
}

// 自增roomId
func (c *ChatRoom) makeRoomId() uint32 {
	once := new(sync.Once)

	var roomId uint32
	once.Do(func() {
		roomId = uint32(util.Redis.Incr(util.Redis.Context(), common.UserIdIncrementKey).Val())
	})

	return roomId
}

func FindOne(roomId uint32) (baseCommon.ChatRoom, interfaces.Error) {
	result := util.Redis.HGet(util.Redis.Context(), common.ChatRoomHashMapKey, strconv.Itoa(int(roomId)))
	if result.Err() != nil {
		return baseCommon.ChatRoom{}, nil
	}
	var room baseCommon.ChatRoom
	b, _ := result.Bytes()
	_ = json.Unmarshal(b, &room)
	return room, nil
}
