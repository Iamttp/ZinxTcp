package main

import "awesomeProject/myZinx/myNet"

func main() {
	s := myNet.NewServe("my")
	s.Start()
}
