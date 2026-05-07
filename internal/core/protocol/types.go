package protocol

import (
	"net"
	"time"
)

const Version uint16 = 4
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

	JoinCodeRequest
	JoinCodeResponse

	ListClientsRequest
	ListClientsResponse

	RequestConnection //
	RequestCodeConnection
	AcceptConnection //
	RejectConnection //

)

type Packet struct {
	Type     PacketType
	SenderID uint32
	TargetID uint32
	Data     []byte
}

type Client struct {
	ID        uint32
	Conn      net.Conn
	SendQueue chan Packet

	Visible  bool
	JoinCode string

	LastSeen time.Time
}
