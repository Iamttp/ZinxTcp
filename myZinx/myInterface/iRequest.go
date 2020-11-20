package myInterface

type IRequest interface {
	GetData() []byte

	GetCnt() int

	GetConnect() IConnect
}
