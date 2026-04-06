package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func dispatch(s *Server, c *protocol.Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketPing:
		return handlePing(c)

	case protocol.PacketFileStart:
		return handleFileStart(s, c, packet)
	case protocol.PacketFileChunk:
		return handleFileChunk(s, c, packet)
	case protocol.PacketFileClose:
		return handleFileClose(s, c, packet)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
