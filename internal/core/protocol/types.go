package protocol

type PacketType uint8

const (
	PacketEcho PacketType = 1
)

type Packet struct {
	Type PacketType
	Data []byte
}
