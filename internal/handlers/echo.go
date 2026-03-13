package handlers

import (
	"kaoluma/internal/core/protocol"
	"net"
)

func handleEcho(conn net.Conn, packet protocol.Packet) error {
	encoded := protocol.EncodePacket(packet)
	_, err := conn.Write(encoded)
	return err
}
