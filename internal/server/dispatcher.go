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
		return handlers.HandleFileStart(c)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
