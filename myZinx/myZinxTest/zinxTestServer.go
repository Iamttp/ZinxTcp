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
	log.Println("Start Ping Handle")

	msg := request.GetMsg()
	log.Println("Server Get Msg : ", msg)

	data := msg.GetData()
	dataLen := int(msg.GetLen())
	for i := 0; i < dataLen; i++ {
		data[i]--
	}
	msg.SetData(data)
	msg.SetId(2)
	request.GetConnect().SendMsg(msg)
}

type HelloRouter struct {
	myNet.BaseRouter
}

func (pr *HelloRouter) Handle(request myInterface.IRequest) {
	log.Println("Start Hello Handle")

	msg := request.GetMsg()
	log.Println("Server Get Msg : ", msg)

	data := msg.GetData()
	dataLen := int(msg.GetLen())
	for i := 0; i < dataLen; i++ {
		data[i]++
	}
	msg.SetData(data)
	msg.SetId(3)
	request.GetConnect().SendMsg(msg)
}

func main() {
	s := myNet.NewServe()

	pr := &PingRouter{}
	hr := &HelloRouter{}
	s.AddRouter(0, pr)
	s.AddRouter(1, hr)
	s.Serve()
}
