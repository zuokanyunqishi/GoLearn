package util

import "GoLearn/chat/interfaces"

type Request struct {
	//当前链接
	conn interfaces.Connection
	//客户端请求数据
	msg       interfaces.Message
	requestId uint32 //当前请求的id自增数
}

func (r *Request) GetRequestId() uint32 {
	return r.requestId
}

func (r *Request) GetConn() interfaces.Connection {
	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求消息的Id
func (r *Request) GetMsgType() uint32 {
	return r.msg.GetMsgType()
}
