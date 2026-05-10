package client

import (
	"fmt"
	"kaoluma/internal/core/protocol"
)

func dispatch(c *Client, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketPong:
		return handlePong(c, packet)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
