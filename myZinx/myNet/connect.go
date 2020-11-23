package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"awesomeProject/myZinx/untils"
	"log"
	"net"
)

type Connect struct {
	Conn     *net.TCPConn
	ConnID   uint32
	IsClosed bool
	ExitChan chan bool
	router   myInterface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, call myInterface.IRouter) *Connect {
	c := &Connect{
		Conn:     conn,
		ConnID:   connID,
		IsClosed: false,
		router:   call,
		ExitChan: make(chan bool, 1), // TODO
	}
	return c
}

func (c *Connect) StartRead() {
	defer c.Stop()

	dpk := NewDataPack()

	for {
		buf := make([]byte, untils.GlobalObj.MaxReadSize) // max 512 byte
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			log.Println("Read error ", err)
			break
		}

		msg, err := dpk.Unpack(buf[:cnt])
		if err != nil {
			log.Println("Read Data Error ", err)
			continue
		}
		// TODO
		r := &Request{
			conn: c,
			msg:  msg,
		}
		c.router.PreHandle(r)
		c.router.Handle(r)
		c.router.PostHandle(r)
	}
}

func (c *Connect) Start() {
	log.Println("Start iConnect ID:", c.ConnID)
	go c.StartRead()
}

func (c *Connect) Stop() {
	log.Println("Stop iConnect ID:", c.ConnID)
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

func (c *Connect) SendMsg(msg myInterface.IMessage) {
	dpk := NewDataPack()
	binary, err := dpk.Pack(msg)
	if err != nil {
		log.Println("Msg Pack Error ", err)
		return
	}

	_, err = c.Conn.Write(binary)
	if err != nil {
		log.Println("Send Msg ", err)
		return
	}
}
