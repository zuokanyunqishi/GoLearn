package interfaces

type Server interface {
	// Start 启动服务
	Start()
	// Stop 停止服务
	Stop()
	// Serve 启动serv
	Serve()

	// AddRouter 添加路由,给当前服务注册路由,什么消息id,对应什么router
	AddRouter(uint32, Router)

	// SetOnConnStart 注册OnConnStart 钩子函数
	SetOnConnStart(func(connection Connection))

	// SetOnConnStop 注册OnConnStop 钩子函数
	SetOnConnStop(func(connection Connection))

	// CallOnConnStart 调用 OnConnStart 钩子函数
	CallOnConnStart(connection Connection)

	// CallOnConnStop 调用 OnConnStop 钩子函数
	CallOnConnStop(connection Connection)
}
