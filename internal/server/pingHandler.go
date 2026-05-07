package server

import (
	"kaoluma/internal/core/protocol"
)

func handlePing(c *ClientInfo) error {
	pong := protocol.Packet{
		Type: protocol.PacketPong,
	}
	c.SendQueue <- pong
	return nil
}
