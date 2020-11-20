package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Client Start...")

	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Fatal("Client Start Error ", err)
	}

	for {
		conn.Write([]byte("hello"))
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Read Error ", err)
			break
		}
		log.Println(string(buf[:n]))
		time.Sleep(time.Second)
	}
}
