package protocol

import (
	"net"
)

const Version uint16 = 1

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
	Conn net.Conn

	//fileState os.File
}

type FileStart struct {
	nameLen  uint16
	name     []byte
	fileSize uint64
}
