package myNet

import (
	"log"
	"net"
	"strconv"
)

type ServeNode struct {
	Name      string
	Ip        string
	Port      int
	IpVersion string
}

func (s *ServeNode) Serve() {
	s.Start()

}

func Handle(conn *net.TCPConn, data []byte, cnt int) error {
	for i := 0; i < cnt; i++ {
		data[i]++
	}

	_, err := conn.Write(data[:cnt])
	return err
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

		c := NewConnection(conn, connID, Handle)
		connID++
		go c.Start()
	}
}

func (s *ServeNode) Stop() {

}

func NewServe(nameServe string) *ServeNode {
	s := &ServeNode{
		Name:      nameServe,
		Ip:        "0.0.0.0",
		Port:      8999,
		IpVersion: "tcp4",
	}
	return s
}
