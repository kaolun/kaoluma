package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func dispatch(s *Server, c *protocol.Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketPing:
		return handlePing(c)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
