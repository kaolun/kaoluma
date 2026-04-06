package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/server/handlers"
)

func dispatch(c *protocol.Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketPing:
		return handlers.HandlePing(c)

	case protocol.PacketFileStart:
		return handlers.HandleFileStart(c, packet)
	case protocol.PacketFileChunk:
		return handlers.HandleFileChunk(c, packet)
	case protocol.PacketFileClose:
		return handlers.HandleFileClose(c, packet)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
