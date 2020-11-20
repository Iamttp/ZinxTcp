package myInterface

type IServer interface {
	Serve()

	Stop()

	Start()

	SetRouter(router IRouter)
}
