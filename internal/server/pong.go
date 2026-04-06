package server

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
)

func handlePing(c *protocol.Client) error {
	err := transport.SendPacket(c.Conn, protocol.Packet{
		Type: protocol.PacketPong,
	})
	if err != nil {
		return err
	}
	return nil
}
