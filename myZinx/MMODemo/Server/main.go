package main

import (
	"awesomeProject/myZinx/MMODemo/Server/Person"
	"awesomeProject/myZinx/MMODemo/Server/core"
	"awesomeProject/myZinx/MMODemo/Server/util"
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
		if play.Pid == player.Pid {
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
		play.SyncUnPid(4, id)
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
	jsonRecv := core.Json3{}
	//log.Println(string(data))
	json.Unmarshal(data, &jsonRecv)

	//player := wm.GetPlayerById(jsonRecv.Id)
	//log.Println(jsonRecv.Id, player)
	player.Pos.X = jsonRecv.X
	player.Pos.Y = jsonRecv.Y
	player.IPerson.SetState(jsonRecv.State)
	player.IPerson.SetMoveVec(&util.Vector2{
		X: jsonRecv.MoveVecX,
		Y: jsonRecv.MoveVecY,
	})
	players := wm.GetAllPlayers()
	if jsonRecv.State == Person.Attack {
		log.Println(player.Pid, " Attack")
	}
	for _, play := range players {
		if play.Pid == player.Pid {
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
