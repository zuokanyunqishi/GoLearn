package interfaces

type DataPack interface {
	// GetHeadLen 获取包头长度
	GetHeadLen() uint32
	// Pack 封装包体
	Pack(Message) ([]byte, error)
	// UnPack 解包
	UnPack([]byte) (Message, error)
}
