package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"log"
	"net"
)

type Connect struct {
	Conn      *net.TCPConn
	ConnID    uint32
	IsClosed  bool
	HandleAPI myInterface.HandleFunc
	ExitChan  chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, call myInterface.HandleFunc) *Connect {
	c := &Connect{
		Conn:      conn,
		ConnID:    connID,
		IsClosed:  false,
		HandleAPI: call,
		ExitChan:  make(chan bool, 1), // TODO
	}
	return c
}

func (c *Connect) StartRead() {
	defer c.Stop()

	for {
		buf := make([]byte, 512) // max 512 byte
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("Read error ", err)
			break
		}

		err = c.HandleAPI(c.Conn, buf, cnt)
		if err != nil {
			log.Println("Handle API error ", err)
		}
	}
}

func (c *Connect) Start() {
	log.Println("Start Connect ID:", c.ConnID)
	go c.StartRead()
}

func (c *Connect) Stop() {
	log.Println("Stop Connect ID:", c.ConnID)
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connect) GetTcpConnect() *net.TCPConn {
	return c.Conn
}

func (c *Connect) GetIdConnect() uint32 {
	return c.ConnID
}

func (c *Connect) GetRemoteAdd() net.Addr {
	return c.Conn.RemoteAddr()
}
