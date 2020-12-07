package main

import (
	"awesomeProject/myZinx/MMODemo/Server/core"
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/myNet"
	"encoding/json"
	"log"
)

func onConnStart(conn myInterface.IConnect) {
	log.Println(conn.GetIdConnect(), "上线")
	player := core.NewPlayer(conn)
	player.SyncPid()
	player.SyncPos()

	players := wm.GetAllPlayers()
	for _, play := range players {
		play.SyncOtherPos(player.Pid, player.Pos.X, player.Pos.Y)
	}
	wm.Add(player)
}

func onConnStop(conn myInterface.IConnect) {
	log.Println(conn.GetIdConnect(), "下线")
	wm.Remove(int32(conn.GetIdConnect())) // TODO getIdConnect 和 player pid 相等 隐患
}

type moveRouter struct {
	myNet.BaseRouter
}

func (pr *moveRouter) Handle(request myInterface.IRequest) {
	log.Println("Start Move Handle")

	player := wm.GetPlayerById(int32(request.GetConnect().GetIdConnect()))
	msg := request.GetMsg()
	data := msg.GetData()
	json.Unmarshal(data, &player.Pos)

	log.Println("Server Get Msg : ", player.Pos)

	players := wm.GetAllPlayers()
	for _, play := range players {
		play.SyncOtherPos(player.Pid, player.Pos.X, player.Pos.Y)
	}
}

var wm *core.WorldManager

func main() {
	s := myNet.NewServe()
	wm = core.NewWorldManager()

	pr := &moveRouter{}
	s.AddRouter(201, pr) // 201 移动请求

	s.SetOnConnStart(onConnStart)
	s.SetOnConnStop(onConnStop)
	s.Serve()
}
