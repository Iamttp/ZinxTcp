package myInterface

import "net"

type Connect interface {
	Start()

	Stop()

	GetTcpConnect() *net.TCPConn

	GetIdConnect() uint32

	GetRemoteAdd() net.Addr
}

type HandleFunc func(*net.TCPConn, []byte, int) error
