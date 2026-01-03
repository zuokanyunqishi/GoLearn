package util

import (
	"GoLearn/chat/util/zlog"
	"errors"
	"io"
	"net"
	"sync"
)

type Connection struct {
	//当前链接socket TCP套接字
	Conn *net.TCPConn
	//链接id
	ConnId uint32
	//当期链接状态
	IsClosed bool
	//退出的channel
	ExitChan chan bool
	//读向写协程传递数据的chan
	MsgChan chan []byte

	//当前链接属于哪个server
	TcpServer *Server

	// 当前连接属于哪个client
	TcpClient *Client

	//链接属性
	property map[string]interface{}

	//属性变更的锁
	propertyLock sync.RWMutex
}

// 设置锁
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()

	defer c.propertyLock.Unlock()
	c.property[key] = value
}

// 获得属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, errors.New("not found key :" + key)
}

// 移除属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if _, ok := c.property[key]; ok {
		delete(c.property, key)
	}
}

// 启动链接
func (c *Connection) Start() {
	zlog.Infof("conn start.... connId  %d", c.ConnId)

	go c.StartReader()

	go c.StartWriter()

	if c.TcpClient != nil {
		c.TcpClient.CallOnConnStart(c)
	}

	if c.TcpServer != nil {
		c.TcpServer.CallOnConnStart(c)
	}
}

// 读取任务携程
func (c *Connection) StartReader() {

	zlog.Infof(
		"[reader Gorouting is running] connId 【%d】 , addr 【%s】 ", c.ConnId, c.RemoteAddr().String())
	defer zlog.Infof(" [链接 connId = %d 关闭....]  ", c.GetConnId())
	defer c.Stop()

	var requestId uint32
	for {

		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTcpConnection(), headData)
		if err != nil {
			zlog.Errorf("read msg error %s", err.Error())
			break
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			zlog.Infof("unpack err %s", err.Error())
			break
		}

		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				zlog.Infof("read msg data error %s", err.Error())
				break
			}

		}
		msg.SetData(data)
		requestId += 1
		req := Request{
			conn:      c,
			msg:       msg,
			requestId: requestId,
		}

		if c.TcpClient != nil {
			c.TcpClient.DoTcpClientRequest(&req)
		}

		if c.TcpServer != nil {
			c.TcpServer.DoTcpServerRequest(&req)
		}

	}

}

// 停止服务
func (c *Connection) Stop() {
	zlog.Warnf("conn stop connId %d ", c.ConnId)
	//链接已经关闭
	if c.IsClosed {
		return
	}
	defer close(c.MsgChan)
	defer close(c.ExitChan)
	defer c.Conn.Close()

	//调用server停止后的 func
	defer func() {
		if c.TcpClient != nil {
			c.TcpClient.CallOnConnStop(c)
		}
		if c.TcpServer != nil {
			c.TcpServer.CallOnConnStop(c)
		}
	}()
	c.IsClosed = true
	c.ExitChan <- true
	//删除链接

	if c.TcpServer != nil {
		c.TcpServer.connManager.RemoveConn(c)
	}

	if c.TcpClient != nil {
		c.TcpClient.connManager.RemoveConn(c)
	}

}

// 获取当前链接
func (c *Connection) GetTcpConnection() *net.TCPConn {

	return c.Conn
}

// 获取当前链接远程地址
func (c *Connection) RemoteAddr() net.Addr {

	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) SendMsg(msgId uint32, data []byte) error {

	if c.IsClosed {
		return errors.New("connection closed when send msg")
	}
	//将数据封包
	dp := NewDataPack()
	bmsg, err := dp.Pack(NewMessage(msgId, data)) //二进制数据
	if err != nil {
		return errors.New("pack data error")
	}

	if _, err := c.GetTcpConnection().Write(bmsg); err != nil {
		return errors.New("send msg error")
	}

	return nil
}

// 获取当前链接id
func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

// 写数据协程
func (c *Connection) StartWriter() {
	zlog.Info("[ writer Goroutine running .......... ]")
	defer zlog.Infof(
		"[ writer Goroutineing  write exit , connid 【 %d 】  RemoteAddr 【 %s】", c.ConnId, c.RemoteAddr().String())
	//不断的阻塞写数据
	for {
		select {
		case data := <-c.MsgChan:
			//拿到了写的消息
			if _, err := c.GetTcpConnection().Write(data); err != nil {
				zlog.Infof(" write msg error %s ", err.Error())
				return
			}
		case <-c.ExitChan:
			//链接已经关闭
			return
		}

	}

}

func (c *Connection) SetTcpServer(s *Server) {
	c.TcpServer = s
}

func (c *Connection) SetTcpClient(client *Client) {
	c.TcpClient = client
}

func NewConnection(conn *net.TCPConn, connId uint32) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connId,
		IsClosed: false,
		ExitChan: make(chan bool),
		MsgChan:  make(chan []byte),
		property: make(map[string]interface{}),
	}

	//添加链接至连接管理map
	return c
}
