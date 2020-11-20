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

	cnt := request.GetCnt()
	data := request.GetData()
	cnn := request.GetConnect().GetTcpConnect()

	for i := 0; i < cnt; i++ {
		data[i]--
	}

	_, err := cnn.Write(data[:cnt])
	if err != nil {
		log.Println("Write Error ")
		return
	}
}

func main() {
	s := myNet.NewServe("my")

	pr := &PingRouter{}
	s.SetRouter(pr)
	s.Start()
}
