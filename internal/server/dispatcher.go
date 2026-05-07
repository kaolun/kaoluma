package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func dispatch(s *Server, c *ClientInfo, packet protocol.Packet) error {
	if packet.TargetID == 0 {
		return handleServerPacket(s, c, packet)
	}

	if isBypassPacket(packet.Type) {
		return s.forward(c, packet)
	}

	if !s.allowed(c.ID, packet.TargetID) {
		return fmt.Errorf("client communication Blocked")
	}

	return s.forward(c, packet)
}

func handleServerPacket(s *Server, c *ClientInfo, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketPing:
		return handlePing(c)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
