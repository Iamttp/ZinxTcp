package main

import (
	"awesomeProject/myZinx/MMODemo/Server/core"
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/myNet"
	"log"
)

func onConnStart(conn myInterface.IConnect) {
	log.Println(conn.GetIdConnect(), "上线")
	player := core.NewPlayer(conn)
	player.SyncPid()
	player.SyncPos()
}

func onConnStop(conn myInterface.IConnect) {
	log.Println(conn.GetIdConnect(), "下线")
}

func main() {
	s := myNet.NewServe()

	//pr := &PingRouter{}
	//hr := &HelloRouter{}
	//s.AddRouter(0, pr)
	//s.AddRouter(1, hr)

	s.SetOnConnStart(onConnStart)
	s.SetOnConnStop(onConnStop)
	s.Serve()
}
