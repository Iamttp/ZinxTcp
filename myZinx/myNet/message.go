package myNet

type Message struct {
	len  uint32
	id   uint32
	data []byte
}

func (m *Message) SetLen(len uint32) {
	m.len = len
}

func (m *Message) SetId(id uint32) {
	m.id = id
}

func (m *Message) SetData(data []byte) {
	m.data = data
}

func (m *Message) GetLen() uint32 {
	return m.len
}

func (m *Message) GetId() uint32 {
	return m.id
}

func (m *Message) GetData() []byte {
	return m.data
}
