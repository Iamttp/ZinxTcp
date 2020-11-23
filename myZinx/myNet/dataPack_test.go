package myNet

import (
	"log"
	"testing"
)

func TestDataPack_Pack(t *testing.T) {
	dpk := NewDataPack()

	res0 := Message{
		len:  2,
		id:   1,
		data: make([]byte, 2),
	}
	res0.data[0] = 1
	res0.data[1] = 2

	res1, err := dpk.Pack(&res0)
	if err != nil {
		log.Println(err)
		return
	}
	res2, err := dpk.Unpack(res1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("push message :", res0)
	log.Println("packed and pop :", res1)
	log.Println("unpacked and pop :", res2)
}
