package client

import (
	"fmt"
	"kaoluma/internal/client/handlers"
	"kaoluma/internal/core/protocol"
)

func dispatch(c *protocol.Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketPong:
		return handlers.HandlePong(packet)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
