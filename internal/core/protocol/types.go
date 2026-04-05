package protocol

import "net"

const Version uint16 = 2

type PacketType uint8

const (
	PacketHandshake PacketType = iota + 1
	PacketHandshakeAck
	PacketHandshakeReject

	PacketEcho
	PacketPing
	PacketPong

	PacketFileStart
	PacketFileChunk
	PacketFileClose
)

type Packet struct {
	Type PacketType
	Data []byte
}

type Client struct {
	Conn      net.Conn
	SendQueue chan Packet
}
