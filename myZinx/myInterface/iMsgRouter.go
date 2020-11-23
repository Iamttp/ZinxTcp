package myInterface

type IMsgRouter interface {
	AddRouter(id uint32, router IRouter)

	DoMsgHandier(request IRequest)
}
