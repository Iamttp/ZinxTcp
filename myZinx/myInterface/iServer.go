package myInterface

type IServer interface {
	Serve()

	Stop()

	Start()

	AddRouter(id uint32, router IRouter)

	GetManager() IConnManager

	SetOnConnStart(onConnStart func(conn IConnect))

	SetOnConnStop(onConnStop func(conn IConnect))

	CallOnConnStart(conn IConnect)

	CallOnConnStop(conn IConnect)
}
