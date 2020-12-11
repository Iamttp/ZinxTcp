package core

import (
	"awesomeProject/myZinx/MMODemo/Server/Person"
	"awesomeProject/myZinx/MMODemo/Server/util"
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/myNet"
	"encoding/json"
	"log"
	"math/rand"
	"sync"
)

type Player struct {
	Pid  int32                // 玩家id，数据库主键生成
	Conn myInterface.IConnect // 玩家连接

	Pos    *util.Vector2  // 玩家位置
	Person Person.IPerson // 玩家操作人物
}

func (p *Player) SendMsg(msgId uint32, data []byte) {
	msg := myNet.Message{}
	msg.SetId(msgId)
	msg.SetData(data)
	msg.SetLen((uint32)(len(data)))
	p.Conn.SendMsg(&msg)
}

type json1 struct {
	Id int32
}

func (p *Player) SyncPid() {
	log.Println("SyncPid")
	data, err := json.Marshal(json1{Id: p.Pid}) // TODO json
	if err != nil {
		log.Println(err)
		return
	}
	p.SendMsg(1, data)
}

func (p *Player) SyncPos() {
	log.Println("SyncPos")
	data, err := json.Marshal(p.Pos) // TODO json
	if err != nil {
		log.Println(err)
		return
	}
	p.SendMsg(2, data)
}

type json3 struct {
	Id int32
	X  float32
	Y  float32
}

func (p *Player) SyncOtherPos(id int32, X float32, Y float32) {
	//log.Println("SyncOtherPos")
	data, err := json.Marshal(json3{
		Id: id,
		X:  X,
		Y:  Y,
	})
	if err != nil {
		log.Println(err)
		return
	}
	p.SendMsg(3, data)
}

func (p *Player) SyncUnPid(msgid uint32, pid int32) {
	log.Println("SyncPid")
	data, err := json.Marshal(json1{Id: pid}) // TODO json
	if err != nil {
		log.Println(err)
		return
	}
	p.SendMsg(msgid, data)
}

//////////////////////////////////////////////////////////////////////

var pidGen int32 = 0
var pidLock sync.Mutex

func NewPlayer(conn myInterface.IConnect) *Player {
	var vec util.Vector2

	pidLock.Lock()
	id := pidGen
	pidGen++
	vec.X = -4 + rand.Float32()
	vec.Y = 2 + rand.Float32()
	pidLock.Unlock()

	return &Player{
		Pid:    id,
		Conn:   conn,
		Pos:    &vec,
		Person: &Person.Person{},
	}
}
