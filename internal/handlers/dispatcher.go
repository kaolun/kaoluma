package handlers

import (
	"fmt"
	"kaoluma/internal/core/protocol"
	"net"
)

func Dispatch(conn net.Conn, packet protocol.Packet) error {
	switch packet.Type {

	case protocol.PacketEcho:
		return handleEcho(conn, packet)

	case protocol.PacketPing:
		return Handlepong(conn)

	default:
		return fmt.Errorf("unknown packet type %d", packet.Type)
	}
}
