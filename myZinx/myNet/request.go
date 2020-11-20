package myNet

import "awesomeProject/myZinx/myInterface"

type Request struct {
	conn myInterface.IConnect
	data []byte
	cnt  int
}

func (r *Request) GetData() []byte {
	return r.data
}

func (r *Request) GetConnect() myInterface.IConnect {
	return r.conn
}

func (r *Request) GetCnt() int {
	return r.cnt
}
