package main

import (
	"awesomeProject/myZinx/myNet"
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Client Start...")

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Fatal("Client Start Error ", err)
	}

	data := [...]byte{1, 2, 3}

	sendMsg := myNet.Message{}
	sendMsg.SetData(data[:])
	sendMsg.SetId(0)
	sendMsg.SetLen(uint32(len(data)))

	dpk := myNet.NewDataPack()
	binaryData, err := dpk.Pack(&sendMsg)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		conn.Write(binaryData)

		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Read Error ", err)
			break
		}
		getMsg, err := dpk.Unpack(buf[:n])
		if err != nil {
			log.Println("Read Message Error ", err)
			break
		}

		log.Println("Client Get Msg : ", getMsg)
		time.Sleep(time.Second)
	}
}
