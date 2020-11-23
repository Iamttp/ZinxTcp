package myInterface

type IDataPack interface {
	GetHeadLen() uint32

	Unpack(dataPack []byte) (IMessage, error)

	Pack(message IMessage) ([]byte, error)
}
