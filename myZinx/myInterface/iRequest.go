package myInterface

type IRequest interface {
	GetConnect() IConnect

	GetMsg() IMessage

	GetData() []byte

	GetCnt() uint32

	GetId() uint32
}
