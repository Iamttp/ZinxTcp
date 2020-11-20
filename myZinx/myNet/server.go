package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"log"
	"net"
	"strconv"
)

type ServeNode struct {
	Name      string
	Ip        string
	Port      int
	IpVersion string
	router    myInterface.IRouter
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
		go c.Start()
	}
}

func (s *ServeNode) Stop() {

}

func (s *ServeNode) SetRouter(router myInterface.IRouter) {
	s.router = router
}

func NewServe(nameServe string) *ServeNode {
	s := &ServeNode{
		Name:      nameServe,
		Ip:        "0.0.0.0",
		Port:      8999,
		IpVersion: "tcp4",
		router:    nil,
	}
	return s
}
