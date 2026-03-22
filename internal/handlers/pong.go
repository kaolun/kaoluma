package handlers

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
)

func Handlepong(c *protocol.Client) error {
	err := transport.SendPacket(c.Conn, protocol.Packet{
		Type: protocol.PacketPong,
	})
	if err != nil {
		return err
	}
	return nil
}
