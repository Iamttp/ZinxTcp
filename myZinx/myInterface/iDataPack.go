package myInterface

type IDataPack interface {
	Unpack(dataPack []byte, startIndex int) (IMessage, error)

	Pack(message IMessage) ([]byte, error)
}
