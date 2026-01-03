package util

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/zlog"
	"fmt"
	"net"
)

var connid uint32 // 连接id
type Client struct {
	//服务器名称
	Name string
	//版本
	IPVersion string
	//端口
	RemotePort int

	RemoteHost string
	//当前消息的Message handler
	MsgHandler interfaces.MessageHandle

	//server 链接管理模块
	connManager interfaces.ConnectionManage

	//server创建后的函数
	connStart func(connection interfaces.Connection)

	//server 销毁后的函数
	connStop func(connection interfaces.Connection)
}

func (c *Client) Start() {

	c.MsgHandler.(*MessageHandler).StarWorkPool()
	raddr, err := net.ResolveTCPAddr(c.IPVersion, fmt.Sprintf("%s:%d", c.RemoteHost, c.RemotePort))
	if err != nil {
		zlog.Fatalf("%s", err.Error())
	}
	tcp, err := net.DialTCP(c.IPVersion, nil, raddr)

	if err != nil {
		zlog.PrintFatalf("链接服务器错误 [%s]", err.Error())
	}
	connid += 1
	connection := NewConnection(tcp, connid)
	connection.SetTcpClient(c)
	connection.Start()
	c.connManager.AddConn(connection)
}

func (c *Client) Stop() {
	//服务资源回收
	zlog.Info("回收资源开始......")
	c.connManager.ClearConn()

	zlog.Info("回收资源成功......")
}

func (c *Client) Serve() {
	c.Start()
}

func (c *Client) AddRouter(u uint32, router interfaces.Router) {
	c.MsgHandler.AddRouterMap(u, router)
}

func (c *Client) SetOnConnStart(f func(connection interfaces.Connection)) {
	c.connStart = f
}

func (c *Client) SetOnConnStop(f func(connection interfaces.Connection)) {
	c.connStop = f
}

func (c *Client) CallOnConnStart(connection interfaces.Connection) {
	if c.connStart != nil {
		c.connStart(connection)
	}
}

func (c *Client) CallOnConnStop(connection interfaces.Connection) {
	if c.connStop != nil {
		c.connStop(connection)
	}
}

// 获取链接管理器
func (c *Client) GetConnManager() interfaces.ConnectionManage {
	return c.connManager
}

func (c *Client) CreateConnection(openWorkPool bool) interfaces.Connection {

	if openWorkPool {
		c.MsgHandler.(*MessageHandler).StarWorkPool()
	}

	raddr, err := net.ResolveTCPAddr(c.IPVersion, fmt.Sprintf("%s:%d", c.RemoteHost, c.RemotePort))
	if err != nil {
		zlog.Fatalf("%s", err.Error())
	}
	tcp, err := net.DialTCP(c.IPVersion, nil, raddr)

	if err != nil {
		zlog.PrintFatalf("链接服务器错误 [%s]", err.Error())
	}
	connid += 1
	connection := NewConnection(tcp, connid)
	connection.SetTcpClient(c)
	connection.Start()
	c.connManager.AddConn(connection)
	return connection
}

func (c *Client) DoTcpClientRequest(request interfaces.Request) {
	ptrM := c.MsgHandler.(*MessageHandler)
	if ptrM.workPoolInitialize {
		ptrM.SendMsgToTaskQueue(request)
		zlog.Infof("requestId %d send taskQueue", request.GetRequestId())
		return
	}
	go ptrM.DoMessageHandle(request)

}

func NewClientServer() *Client {
	return &Client{
		Name:        "chatClient",
		IPVersion:   "tcp4",
		RemotePort:  8805,
		RemoteHost:  "127.0.0.1",
		MsgHandler:  NewMessageHandler(),
		connManager: NewConnectionManager(),
	}
}
