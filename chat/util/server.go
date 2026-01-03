package util

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/zlog"
	"fmt"
	"net"
)

type Server struct {
	//服务器名称
	Name string
	//版本
	IPVersion string
	//端口
	Port int

	Host string
	//当前消息的Message handler
	MsgHandler *MessageHandler

	//server 链接管理模块
	connManager interfaces.ConnectionManage

	//server创建后的函数
	connStart func(connection interfaces.Connection)

	//server 销毁后的函数
	connStop func(connection interfaces.Connection)
}

// 注册OnConnStart 钩子函数
func (s *Server) SetOnConnStart(start func(connection interfaces.Connection)) {
	s.connStart = start
}

// 注册OnConnStop 钩子函数
func (s *Server) SetOnConnStop(stop func(connection interfaces.Connection)) {
	s.connStop = stop
}

// 调用 OnConnStart 钩子函数
func (s *Server) CallOnConnStart(connection interfaces.Connection) {
	if s.connStart != nil {
		s.connStart(connection)
	}
}

// 调用 OnConnStop 钩子函数
func (s *Server) CallOnConnStop(connection interfaces.Connection) {

	if s.connStop != nil {
		s.connStop(connection)
	}
}

// 开启服务
func (s *Server) Start() {

	zlog.PrintInfof("server %s Host %s Port %d start ", s.Name, s.Host, s.Port)

	s.MsgHandler.StarWorkPool()
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.Host, s.Port))

	if err != nil {
		zlog.PrintFatalf("%s", err.Error())
		return
	}

	tcp, _ := net.ListenTCP(s.IPVersion, addr)

	var cid uint32
	for {
		conn, err := tcp.AcceptTCP()
		if err != nil {
			zlog.Error(err.Error())
		}

		connection := NewConnection(conn, cid)
		connection.SetTcpServer(s)
		cid += 1

		s.connManager.AddConn(connection)

		go connection.Start()

	}

}

// 停止服务
func (s *Server) Stop() {

	//服务资源回收
	zlog.Info("回收资源开始......")
	s.connManager.ClearConn()

	zlog.Info("回收资源成功......")

}

// 运行服务
func (s *Server) Serve() {

	if len(s.MsgHandler.MsgHandleMap) <= 0 {
		zlog.PrintFatal("路由未设置,终止..")
		return
	}
	s.Start()
	select {}
}

// 添加router
func (s *Server) AddRouter(msgId uint32, router interfaces.Router) {
	s.MsgHandler.AddRouterMap(msgId, router)
}

// 获取链接管理器
func (s *Server) GetConnManager() interfaces.ConnectionManage {
	return s.connManager
}

func (s *Server) DoTcpServerRequest(request interfaces.Request) {

	if s.MsgHandler.workPoolInitialize {
		s.MsgHandler.SendMsgToTaskQueue(request)
		return
	}
	go s.MsgHandler.DoMessageHandle(request)
}

func NewServer() interfaces.Server {

	return &Server{
		Name:        "chantServer",
		IPVersion:   "tcp4",
		Port:        8805,
		Host:        "0.0.0.0",
		MsgHandler:  NewMessageHandler(),
		connManager: NewConnectionManager(),
	}

}
