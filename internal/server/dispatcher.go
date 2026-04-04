package server

import (
	"fmt"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/server/handlers"
)

func dispatch(c *protocol.Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketEcho:
		return handlers.HandleEcho(c, packet)

	case protocol.PacketPing:
		return handlers.HandlePong(c)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
