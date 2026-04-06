package protocol

import (
	"net"
)

const Version uint16 = 3
const ChunkSize = 50 * 1024
const MaxPacketSize = 70 * 1024

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
	ID        uint32
	Conn      net.Conn
	SendQueue chan Packet
}

type FileStart struct {
	FileID     uint32
	FileSize   uint64
	NameLength uint16
	FileName   []byte
}
type FileChunk struct {
	FileID     uint32
	ChuckIndex uint32
	Data       []byte
}
type FileEnd struct {
	FileID uint32
}
type IncomingFile struct {
	FileName  string
	FileSize  uint64
	FileOwner uint32
	Path      string
}
