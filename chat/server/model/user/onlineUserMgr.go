package user

import (
	baseCommon "GoLearn/chat/common"
	"GoLearn/chat/interfaces"
	"GoLearn/chat/server/common"
	"GoLearn/chat/util/errors"
	"GoLearn/chat/util/zlog"
	"sync"
)

var OnlineUser = NewOnlineUserManage()

type OnlineUserManage struct {
	userMgr map[uint32]*Process
	lock    sync.RWMutex
}

func (m *OnlineUserManage) Add(process *Process) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.userMgr[process.User.GetUserId()] = process

}

func (m *OnlineUserManage) FindOne(userId uint32) (*Process, interfaces.Error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	process, ok := m.userMgr[userId]
	if !ok {
		return process, errors.New("current user not online", common.UserNotOnline)
	}
	return process, nil
}

func (m *OnlineUserManage) Get(userIds []uint32) ([]*Process, interfaces.Error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	var onlineUsers = make([]*Process, len(userIds))
	for _, id := range userIds {
		if user, ok := m.userMgr[id]; ok {
			onlineUsers = append(onlineUsers, user)
		}
	}

	if len(onlineUsers) < 1 {
		return onlineUsers, errors.NewNoCode("无在线用户")
	}
	return onlineUsers, nil
}

func (m *OnlineUserManage) Delete(userId uint32) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.userMgr, userId)
	zlog.Warnf("delete %d from userMgr map", userId)

}

func (m *OnlineUserManage) UserCount() int {
	return len(m.userMgr)
}

func (m *OnlineUserManage) GetAll() []baseCommon.BaseUsers {
	var onlineUSers = make([]baseCommon.BaseUsers, 0)
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, process := range m.userMgr {
		onlineUSers = append(onlineUSers, baseCommon.BaseUsers{
			UserName: process.User.UserName,
			UserId:   process.User.UserId,
			Status:   baseCommon.UserOnLine,
			NickName: process.User.NickName,
		})
	}

	return onlineUSers

}

func (m *OnlineUserManage) IsOnline(userId uint32) bool {
	if _, ok := m.userMgr[userId]; ok {
		return true
	}
	return false
}

func NewOnlineUserManage() *OnlineUserManage {
	return &OnlineUserManage{userMgr: make(map[uint32]*Process)}
}
