package server

import "kaoluma/internal/core/protocol"

func isBypassPacket(t protocol.PacketType) bool {
	switch t {

	case protocol.RequestConnection:
		return true
	case protocol.AcceptConnection:
		return true
	case protocol.RejectConnection:
		return true

	default:
		return false
	}
}
