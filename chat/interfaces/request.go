package interfaces

type Request interface {
	// GetConn 获取链接
	GetConn() Connection
	// GetData 获取数据
	GetData() []byte
	GetMsgType() uint32
	GetRequestId() uint32
}
