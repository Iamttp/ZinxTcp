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
	router   myInterface.IMsgRouter
	ExitChan chan bool
	MsgChan  chan []byte
}

func NewConnection(conn *net.TCPConn, connID uint32, call myInterface.IMsgRouter) *Connect {
	c := &Connect{
		Conn:     conn,
		ConnID:   connID,
		IsClosed: false,
		router:   call,
		ExitChan: make(chan bool, 1),
		MsgChan:  make(chan []byte),
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
			log.Println("Unpack Data Error ", err)
			continue
		}

		r := &Request{
			conn: c,
			msg:  msg,
		}

		go c.router.DoMsgHandier(r)
	}
}

func (c *Connect) StartWrite() {
	for {
		select {
		case binary := <-c.MsgChan: // 二进制数据已传入管道
			_, err := c.Conn.Write(binary)
			if err != nil {
				log.Println("Send Msg ", err)
				return
			}
		case <-c.ExitChan: // Write 退出
			log.Println("Write Return")
			return
		}
	}
}

func (c *Connect) Start() {
	log.Println("Start iConnect ID:", c.ConnID)
	go c.StartRead()
	go c.StartWrite()
}

func (c *Connect) Stop() {
	log.Println("Stop iConnect ID:", c.ConnID)
	if c.IsClosed == true {
		return
	}
	c.IsClosed = true
	c.ExitChan <- true

	c.Conn.Close()
	/*
		channel不需要通过close释放资源，只要没有goroutine持有channel，相关资源会自动释放。
		close可以用来通知channel接收者不会再收到数据。所以即使channel中有数据也可以close而不会导致接收者收不到残留的数据。
		有些场景需要关闭通道，例如range遍历通道，如不关闭range遍历会出现死锁。
	*/
	close(c.ExitChan)
	close(c.MsgChan)
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

	c.MsgChan <- binary
}
