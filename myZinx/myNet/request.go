package myNet

import "awesomeProject/myZinx/myInterface"

type Request struct {
	conn myInterface.IConnect
	msg  myInterface.IMessage
}

func (r *Request) GetConnect() myInterface.IConnect {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetCnt() uint32 {
	return r.msg.GetLen()
}

func (r *Request) GetId() uint32 {
	return r.msg.GetId()
}

func (r *Request) GetMsg() myInterface.IMessage {
	return r.msg
}
