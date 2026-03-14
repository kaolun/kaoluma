package network

import (
	"encoding/binary"
	"fmt"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"net"
)

//func InitiateHandshake(conn net.Conn) error {

//}

func PerformHandshake(conn net.Conn) error {
	packet, err := protocol.DecodePacket(conn)
	if err != nil {
		return err
	}
	if packet.Type != protocol.PacketHandshake {
		return fmt.Errorf("invalid handshake packet type")
	}
	if len(packet.Data) < 2 {
		return fmt.Errorf("invalid handshake version length")
	}

	version := binary.BigEndian.Uint16(packet.Data[:2])

	if version != protocol.Version {
		err := transport.SendPacket(conn, protocol.Packet{
			Type: protocol.PacketHandshakeReject,
		})
		if err != nil {
			return fmt.Errorf("failed to send handshake reject")
		}
		return fmt.Errorf("invalid handshake version length")
	}

	err = transport.SendPacket(conn, protocol.Packet{
		Type: protocol.PacketHandshakeAck,
	})
	if err != nil {
		return fmt.Errorf("failed to write handshake ack")
	}
	return nil
}
