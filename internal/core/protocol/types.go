package protocol

const Version uint16 = 5
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

	VisibilityChangeRequest
	VisibilityChangeConfirm

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
