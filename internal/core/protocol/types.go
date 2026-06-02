package protocol

const Version uint16 = 7
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

	RequestPeerConnection // data code
	AcceptPeerConnection  // data id
	RejectPeerConnection
	NotifyPeerConnection // data id

	DisconnectPeerConnection
	ConfirmPeerDisconnect
	NotifyPeerDisconnect
)

type Packet struct {
	Type     PacketType
	SenderID uint32
	TargetID uint32
	Data     []byte
}
