package transport

import (
	"kaoluma/internal/core/protocol"
	"net"
	"time"
)

func SendPacket(conn net.Conn, packet protocol.Packet) error {
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	encoded := protocol.EncodePacket(packet)
	_, err := conn.Write(encoded)
	return err
}
