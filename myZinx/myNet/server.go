package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/untils"
	"log"
	"net"
	"strconv"
)

type ServeNode struct {
	Name        string
	Ip          string
	Port        int
	IpVersion   string
	router      myInterface.IMsgRouter
	connManager myInterface.IConnManager

	// 添加连接开启和结束方法
	onConnStart func(conn myInterface.IConnect)
	onConnStop  func(conn myInterface.IConnect)

	// TODO iConnect中为每个连接 提供添加 可访问属性 map[string]interface{} 让用户可以在onConnStart中设置
}

func (s *ServeNode) Serve() {
	s.Start()
}

func (s *ServeNode) Start() {
	defer s.Stop()
	log.Println("Start...", s.Ip+":"+strconv.Itoa(s.Port))
	addr, err := net.ResolveTCPAddr(s.IpVersion, s.Ip+":"+strconv.Itoa(s.Port))
	if err != nil {
		log.Fatal("Start Error ", err)
		return
	}

	listener, err := net.ListenTCP(s.IpVersion, addr)
	if err != nil {
		log.Fatal("Listen TCP Error", err)
		return
	}

	log.Println("Start Success !", s.Name)

	var connID uint32
	connID = 0

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Accept error", err)
			continue
		}

		if untils.GlobalObj.MaxConnect <= s.GetManager().Len() {
			// TODO conn 返回到达最大连接数提示
			log.Println("Connect is Maxed")
			conn.Close()
			continue
		}

		c := NewConnection(conn, connID, s.router, s)
		connID++
		c.Start() // TODO go 是否有必要添加go
	}
}

func (s *ServeNode) Stop() {
	log.Println("Stop...", s.Ip+":"+strconv.Itoa(s.Port))
	s.connManager.Clear()
}

func (s *ServeNode) AddRouter(id uint32, router myInterface.IRouter) {
	s.router.AddRouter(id, router)
}

func (s *ServeNode) GetManager() myInterface.IConnManager {
	return s.connManager
}

func (s *ServeNode) SetOnConnStart(onConnStart func(conn myInterface.IConnect)) {
	s.onConnStart = onConnStart
}

func (s *ServeNode) SetOnConnStop(onConnStop func(conn myInterface.IConnect)) {
	s.onConnStop = onConnStop
}

func (s *ServeNode) CallOnConnStart(conn myInterface.IConnect) {
	if s.onConnStart != nil {
		s.onConnStart(conn)
	}
}

func (s *ServeNode) CallOnConnStop(conn myInterface.IConnect) {
	if s.onConnStop != nil {
		s.onConnStop(conn)
	}
}

func NewServe() *ServeNode {
	s := &ServeNode{
		Name:        untils.GlobalObj.Name,
		Ip:          untils.GlobalObj.Ip,
		Port:        untils.GlobalObj.Port,
		IpVersion:   untils.GlobalObj.IpVersion,
		router:      NewMsgRouter(),
		connManager: NewConnManage(),
		onConnStart: nil,
		onConnStop:  nil,
	}
	return s
}
