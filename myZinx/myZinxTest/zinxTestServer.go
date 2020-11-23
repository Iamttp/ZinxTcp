package main

import (
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/myNet"
	"log"
)

type PingRouter struct {
	myNet.BaseRouter
}

func (pr *PingRouter) Handle(request myInterface.IRequest) {
	log.Println("Start Handle")

	msg := request.GetMsg()
	log.Println("Server Get Msg : ", msg)

	if msg.GetId() == 0 {
		data := msg.GetData()
		dataLen := int(msg.GetLen())
		for i := 0; i < dataLen; i++ {
			data[i]--
		}
		msg.SetData(data)
		msg.SetId(1)
		request.GetConnect().SendMsg(msg)
	}
}

func main() {
	s := myNet.NewServe()

	pr := &PingRouter{}
	s.SetRouter(pr)
	s.Start()
}
