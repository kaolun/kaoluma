package handlers

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
)

func handleEcho(c *protocol.Client, packet protocol.Packet) error {
	return transport.SendPacket(c.Conn, packet)
}
