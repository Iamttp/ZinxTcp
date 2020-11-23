package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"log"
)

type MsgRouter struct {
	mpIdRouter map[uint32]myInterface.IRouter
}

func (mg *MsgRouter) AddRouter(id uint32, router myInterface.IRouter) {
	mg.mpIdRouter[id] = router
}

func (mg *MsgRouter) DoMsgHandier(request myInterface.IRequest) {
	val, ok := mg.mpIdRouter[request.GetId()]
	if !ok {
		log.Println("Cannot found Id ")
		return
	}

	val.PreHandle(request)
	val.Handle(request)
	val.PostHandle(request)
}

func NewMsgRouter() myInterface.IMsgRouter {
	return &MsgRouter{
		mpIdRouter: make(map[uint32]myInterface.IRouter),
	}
}
