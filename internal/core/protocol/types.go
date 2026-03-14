package protocol

const Version uint16 = 1

type PacketType uint8

const (
	PacketHandshake       PacketType = 1
	PacketHandshakeAck    PacketType = 2
	PacketHandshakeReject PacketType = 3

	PacketEcho PacketType = 10
)

type Packet struct {
	Type PacketType
	Data []byte
}
