package myInterface

type IMessage interface {
	// len  uint32
	// id   uint32
	// data []byte

	SetLen(len uint32)
	SetId(id uint32)
	SetData(data []byte)

	GetLen() uint32
	GetId() uint32
	GetData() []byte
}
