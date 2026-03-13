package handlers

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"net"
)

func handleEcho(conn net.Conn, packet protocol.Packet) error {
	return transport.SendPacket(conn, packet)
}
