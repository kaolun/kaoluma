package handlers

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func Dispatch(c *protocol.Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketEcho:
		return handleEcho(c, packet)

	case protocol.PacketPing:
		return Handlepong(c)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
