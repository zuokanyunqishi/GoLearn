package util

type Message struct {
	MessageType uint32
	DataLen     uint32
	Data        []byte
}

func (m *Message) GetMsgType() uint32 {
	return m.MessageType
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgType(typeId uint32) {
	m.MessageType = typeId
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}

func (m *Message) SetDataLen(len uint32) {
	m.DataLen = len
}

func NewMessage(typeId uint32, data []byte) *Message {
	return &Message{
		MessageType: typeId,
		DataLen:     uint32(len(data)),
		Data:        data,
	}

}

func BuildTextMsg(msgId uint32, data string) *Message {
	return NewMessage(msgId, []byte(data))
}

func BuildBytesMsg(msgId uint32, data []byte) *Message {
	return NewMessage(msgId, data)
}
