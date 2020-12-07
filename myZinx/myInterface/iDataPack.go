package myInterface

type IDataPack interface {
	Unpack(dataPack []byte) (IMessage, error)

	Pack(message IMessage) ([]byte, error)
}
