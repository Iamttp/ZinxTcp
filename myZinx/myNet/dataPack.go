package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"encoding/binary"
	"log"
)

// len id data

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//func (dpk *DataPack) Unpack(dataPack []byte) (myInterface.IMessage, error) {
//	dr := bytes.NewReader(dataPack)
//	m := Message{}
//
//	err := binary.Read(dr, binary.LittleEndian, &m.len)
//	if err != nil {
//		return nil, err
//	}
//	err = binary.Read(dr, binary.LittleEndian, &m.id)
//	if err != nil {
//		return nil, err
//	}
//
//	// TODO 拷贝性能考虑
//	temp := make([]byte, m.len)
//	//if int(m.len) > untils.GlobalObj.MaxReadSize {
//	//	return nil, err
//	//}
//	err = binary.Read(dr, binary.LittleEndian, &temp)
//	if err != nil {
//		return nil, err
//	}
//	m.data = temp // 切片赋值，指向相同 &m.data[0] == &temp[0]
//	return &m, nil
//}

func (dpk *DataPack) Unpack(dataPack []byte, startIndex int) myInterface.IMessage {
	m := Message{}

	arr := make([]byte, 4)
	copy(arr, dataPack[startIndex:])
	m.len = binary.LittleEndian.Uint32(arr)
	if len(dataPack) <= startIndex+4 {
		log.Println("startIndex + 4 >= len(dataPack) ")
		return nil
	}
	copy(arr, dataPack[startIndex+4:])
	m.id = binary.LittleEndian.Uint32(arr)

	// TODO 拷贝性能考虑
	m.data = make([]byte, m.len)
	if len(dataPack) <= startIndex+8 {
		log.Println("startIndex + 8 >= len(dataPack) ")
		return nil
	}
	copy(m.data, dataPack[startIndex+8:])
	return &m
}

func (dpk *DataPack) Pack(message myInterface.IMessage) []byte {
	var res = make([]byte, 8+message.GetLen())
	binary.LittleEndian.PutUint32(res[0:4], message.GetLen())
	binary.LittleEndian.PutUint32(res[4:8], message.GetId())
	copy(res[8:], message.GetData())
	return res
}
