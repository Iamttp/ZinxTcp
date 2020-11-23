package myInterface

import "net"

type IConnect interface {
	Start()

	Stop()

	GetTcpConnect() *net.TCPConn

	GetIdConnect() uint32

	GetRemoteAdd() net.Addr

	SendMsg(msg IMessage)
}
