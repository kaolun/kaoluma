package handlers

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"net"
)

func Handlepong(conn net.Conn) error {
	err := transport.SendPacket(conn, protocol.Packet{
		Type: protocol.PacketPong,
	})
	if err != nil {
		return err
	}
	return nil
}
