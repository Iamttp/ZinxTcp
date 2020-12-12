package main

import (
	"awesomeProject/myZinx/MMODemo/Server/Person"
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
		if play.Json1.Id == player.Json1.Id {
			continue
		}
		log.Println("SyncOtherPos in onConnStart")
		play.SyncOtherPos(player)
		player.SyncOtherPos(play)
	}
	wm.Add(player)
}

func onConnStop(conn myInterface.IConnect) {
	log.Println(conn.GetIdConnect(), "下线")
	id := int32(conn.GetIdConnect())
	wm.Remove(id) // TODO getIdConnect 和 player pid 相等 隐患

	players := wm.GetAllPlayers()
	for _, play := range players {
		play.SyncUnPid(id)
	}
}

type moveRouter struct {
	myNet.BaseRouter
}

func (pr *moveRouter) Handle(request myInterface.IRequest) {
	// log.Println("Start Move Handle")
	player := wm.GetPlayerById(int32(request.GetConnect().GetIdConnect()))
	msg := request.GetMsg()
	data := msg.GetData()
	json.Unmarshal(data, player.Json3)

	players := wm.GetAllPlayers()
	if player.Json3.State == Person.Attack {
		log.Println(player.Json1.Id, " Attack")
	}
	for _, play := range players {
		if play.Json1.Id == player.Json1.Id {
			continue
		}
		play.SyncOtherPos(player)
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
