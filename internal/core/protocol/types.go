package protocol

const Version uint16 = 1

type PacketType uint8

const (
	PacketHandshake PacketType = iota + 1
	PacketHandshakeAck
	PacketHandshakeReject

	PacketEcho
	PacketPing
	PacketPong
)

type Packet struct {
	Type PacketType
	Data []byte
}
