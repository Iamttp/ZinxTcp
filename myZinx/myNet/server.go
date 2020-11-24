package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/untils"
	"log"
	"net"
	"strconv"
)

type ServeNode struct {
	Name      string
	Ip        string
	Port      int
	IpVersion string
	router    myInterface.IMsgRouter
}

func (s *ServeNode) Serve() {
	s.Start()
}

func (s *ServeNode) Start() {
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

		c := NewConnection(conn, connID, s.router)
		connID++
		c.Start() // TODO go
	}
}

func (s *ServeNode) Stop() {

}

func (s *ServeNode) AddRouter(id uint32, router myInterface.IRouter) {
	s.router.AddRouter(id, router)
}

func NewServe() *ServeNode {
	s := &ServeNode{
		Name:      untils.GlobalObj.Name,
		Ip:        untils.GlobalObj.Ip,
		Port:      untils.GlobalObj.Port,
		IpVersion: untils.GlobalObj.IpVersion,
		router:    NewMsgRouter(),
	}
	return s
}
