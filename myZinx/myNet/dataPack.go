package myNet

import (
	"awesomeProject/myZinx/myInterface"
	"bytes"
	"encoding/binary"
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

func (dpk *DataPack) Unpack(dataPack []byte, startIndex int) (myInterface.IMessage, error) {
	m := Message{}

	arr := make([]byte, 4)
	copy(arr, dataPack[startIndex:])
	m.len = binary.LittleEndian.Uint32(arr)
	copy(arr, dataPack[startIndex+4:])
	m.id = binary.LittleEndian.Uint32(arr)

	// TODO 拷贝性能考虑
	m.data = make([]byte, m.len)
	copy(m.data, dataPack[startIndex+8:])
	return &m, nil
}

func (dpk *DataPack) Pack(message myInterface.IMessage) ([]byte, error) {
	dataBuf := new(bytes.Buffer)

	err := binary.Write(dataBuf, binary.LittleEndian, message.GetLen())
	if err != nil {
		return nil, err
	}
	err = binary.Write(dataBuf, binary.LittleEndian, message.GetId())
	if err != nil {
		return nil, err
	}
	err = binary.Write(dataBuf, binary.LittleEndian, message.GetData())
	if err != nil {
		return nil, err
	}
	return dataBuf.Bytes(), nil
}
