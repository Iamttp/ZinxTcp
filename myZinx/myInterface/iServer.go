package myInterface

type IServer interface {
	Serve()

	Stop()

	Start()

	AddRouter(id uint32, router IRouter)

	GetManager() IConnManager
}
