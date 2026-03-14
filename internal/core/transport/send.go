package transport

import (
	"kaoluma/internal/core/protocol"
	"net"
	"time"
)

func SendPacket(conn net.Conn, packet protocol.Packet) error {
	conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
	encoded, err := protocol.EncodePacket(packet)
	if err != nil {
		return err
	}
	_, err = conn.Write(encoded)
	return err
}
