package util

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/zlog"
	"errors"
	"fmt"
	"sync"
)

type ConnectionManager struct {
	connectionsMap map[uint32]interfaces.Connection

	connLock sync.RWMutex
}

func (c *ConnectionManager) AddConn(conn interfaces.Connection) {

	//共享读写锁map
	c.connLock.Lock()
	defer c.connLock.Unlock()

	//添加链接到map
	c.connectionsMap[conn.GetConnId()] = conn

	zlog.Infof("connId %d add to connMap succ , current map len --> %d ", conn.GetConnId(), c.Len())

}

// 删除链接
func (c *ConnectionManager) RemoveConn(conn interfaces.Connection) {
	//添加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if _, ok := c.connectionsMap[conn.GetConnId()]; ok {
		delete(c.connectionsMap, conn.GetConnId())
		zlog.Warnf("connId %d remove to connMap succ, connMap len  %d ", conn.GetConnId(), c.Len())
		return
	}

	zlog.Infof("connId %d not found ", conn.GetConnId())

}

// 获取当前链接
func (c *ConnectionManager) GetConn(connId uint32) (interfaces.Connection, error) {
	//添加读锁getMap
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	if conn, ok := c.connectionsMap[connId]; ok {
		zlog.Infof("connId %d found ", conn.GetConnId())
		return conn, nil
	}
	return nil, errors.New(fmt.Sprintf("connId %d not found", connId))
}

// 获取链接数
func (c *ConnectionManager) Len() int {
	return len(c.connectionsMap)
}

// 清理链接
func (c *ConnectionManager) ClearConn() {
	//添加写锁
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connId, conn := range c.connectionsMap {
		//停止链接
		conn.Stop()
		//移除链接
		delete(c.connectionsMap, connId)
	}
	zlog.Infof("connMap clear ok,connNum %d", c.Len())

}

func NewConnectionManager() *ConnectionManager {

	return &ConnectionManager{
		connectionsMap: make(map[uint32]interfaces.Connection),
		connLock:       sync.RWMutex{},
	}
}
