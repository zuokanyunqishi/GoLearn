package util

import (
	"GoLearn/chat/interfaces"
	"GoLearn/chat/util/zlog"
)

type MessageHandler struct {
	MsgHandleMap       map[uint32]interfaces.Router //路由map存放messageId => router
	WorkPoolSize       uint32                       //worker 协程池的数量
	TaskWorkQueue      []chan interfaces.Request    //取worker投递的消息chain
	workPoolInitialize bool
	MaxTaskQueueLen    uint32 // 任务队列最大长度
}

// AddRouterMap 添加路由到map集合 key 为msgId ,value 为 Router
func (m *MessageHandler) AddRouterMap(msgType uint32, router interfaces.Router) {
	if _, ok := m.MsgHandleMap[msgType]; ok {
		zlog.Infof("current router exits %d", msgType)
		return
	}
	m.MsgHandleMap[msgType] = router
	zlog.Infof("add router to routerMap success %d", msgType)
}

func (m *MessageHandler) DoMessageHandle(request interfaces.Request) {
	handler, ok := m.MsgHandleMap[request.GetMsgType()]
	if !ok {
		zlog.Infof("api router not found ,must reg")
		return
	}
	//处理消息路由
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PreHandle(request)
}

// StarWorkPool StartWorkerPool() 方法是启动Worker工作池， 这里根据用户配置好
// 的 WorkerPoolSize 的数量来启动， 然后分别给每个Worker分配一
// 个 TaskQueue ， 然后用一个goroutine来承载一个Worker的工作业务。
func (m *MessageHandler) StarWorkPool() {
	if m.workPoolInitialize {
		zlog.Info("work pool is running..")
		return
	}

	if m.WorkPoolSize < 1 {
		zlog.Warnf("work poolSize < 1 , stop init work pool  ")
		return
	}
	for i := 0; i < int(m.WorkPoolSize); i++ {
		m.TaskWorkQueue[i] = make(chan interfaces.Request, m.MaxTaskQueueLen)
		//启动一个work 协程
		go m.startOneWork(i, m.TaskWorkQueue[i])
	}
	m.workPoolInitialize = true
	zlog.Infof("work pool init success ,WorkPoolSize is %d ", m.WorkPoolSize)
}

// 启动一个work 进行任务处理
// StartOneWorker() 方法就是一个Worker的工作业务， 每个worker是不会退出的
// (目前没有设定worker的停止工作机制)， 会永久的从对应的TaskQueue中等待消
// 息， 并处理
func (m *MessageHandler) startOneWork(workId int, taskQueue chan interfaces.Request) {

	for {
		select {
		case request := <-taskQueue:
			m.DoMessageHandle(request)
			zlog.Infof("workId [%d ] handel over , current link requestId [%d] ", workId, request.GetRequestId())
		}

	}

}

// SendMsgToTaskQueue 将消息投递给某个worker进行处理
func (m *MessageHandler) SendMsgToTaskQueue(request interfaces.Request) {

	//StartOneWorker() 方法就是一个Worker的工作业务， 每个worker是不会退出的
	//(目前没有设定worker的停止工作机制)， 会永久的从对应的TaskQueue中等待消
	//息， 并处理
	workId := request.GetRequestId() % m.WorkPoolSize

	zlog.Infof(" ConnId [%d]  link requestId [%d] Send to [%d] worker handel  ",
		request.GetConn().GetConnId(), request.GetRequestId(), workId)
	m.TaskWorkQueue[int(workId)] <- request
}

func (m *MessageHandler) WorkPoolIsInit() bool {
	return m.workPoolInitialize
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		MsgHandleMap:       make(map[uint32]interfaces.Router),
		WorkPoolSize:       30,
		TaskWorkQueue:      make([]chan interfaces.Request, 100),
		workPoolInitialize: false,
	}
}
