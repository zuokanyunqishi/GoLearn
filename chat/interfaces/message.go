package interfaces

type Message interface {
	GetMsgType() uint32 //获取消息typeId
	GetMsgLen() uint32  //获取消息长度
	GetData() []byte    //获取消息内容

	SetMsgType(uint32) //设置消息typeId
	SetData([]byte)    //设置消息内容
	SetDataLen(uint32) //设置消息长度
}
