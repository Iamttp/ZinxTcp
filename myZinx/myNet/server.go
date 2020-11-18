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
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println("Accept error", err)
			continue
		}

		go func() {
			for {
				buf := make([]byte, 512) // max 512 byte
				cnt, err := conn.Read(buf)
				if err != nil {
					log.Println("Read error ", err)
					break
				}

				for i := 0; i < cnt; i++ {
					buf[i]--
				}

				_, err = conn.Write(buf[:cnt])
				if err != nil {
					log.Println("Write error ", err)
					continue
				}
			}
		}()
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
