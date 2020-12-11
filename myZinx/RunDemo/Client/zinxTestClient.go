package main

import (
	"awesomeProject/myZinx/myNet"
	"awesomeProject/myZinx/untils"
	"log"
	"net"
	"time"
)

// Window下先go build得到exe文件，然后新建.bat文件，复制多份start /min ./zinxTestClient.exe，创建多个客户端
func main() {
	log.Println("Client Start...")

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Fatal("Client Start Error ", err)
	}

	data := [...]byte{1, 2, 3}

	sendMsg := myNet.Message{}
	sendMsg.SetData(data[:])
	sendMsg.SetLen(uint32(len(data)))

	dpk := myNet.NewDataPack()

	var i bool
	for {
		if i {
			sendMsg.SetId(1)
		} else {
			sendMsg.SetId(0)
		}
		i = !i
		binaryData := dpk.Pack(&sendMsg)

		conn.Write(binaryData)

		buf := make([]byte, untils.GlobalObj.MaxReadSize)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Read Error ", err)
			break
		}
		getMsg := dpk.Unpack(buf[:n], 0)

		log.Println("Client Get Msg : ", getMsg)
		time.Sleep(time.Second)
	}
}
