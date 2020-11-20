package myNet

import "awesomeProject/myZinx/myInterface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(request myInterface.IRequest)  {}
func (br *BaseRouter) Handle(request myInterface.IRequest)     {}
func (br *BaseRouter) PostHandle(request myInterface.IRequest) {}
