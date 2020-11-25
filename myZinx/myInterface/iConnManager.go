package myInterface

type IConnManager interface {
	Get(index uint32) IConnect
	Len() int
	Add(connect IConnect)
	Remove(connect IConnect)
	Clear()
}
