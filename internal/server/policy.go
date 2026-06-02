package server

import "kaoluma/internal/core/protocol"

func isBypassPacket(t protocol.PacketType) bool {
	switch t {

	default:
		return false
	}
}
