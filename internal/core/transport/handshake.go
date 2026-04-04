package transport

import (
	"encoding/binary"
	"fmt"
	"kaoluma/internal/core/protocol"
	"net"
)

func InitiateHandshake(conn net.Conn) error {
	versionBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(versionBuf, protocol.Version)

	packet := protocol.Packet{
		Type: protocol.PacketHandshake,
		Data: versionBuf,
	}

	err := SendPacket(conn, packet)
	if err != nil {
		return fmt.Errorf("failed to send handshake: %d", err)
	}

	responsePacket, err := protocol.DecodePacket(conn)
	if err != nil {
		return fmt.Errorf("failed to decode handshake response: %d", err)
	}

	switch responsePacket.Type {
	case protocol.PacketHandshakeAck:
		return nil
	case protocol.PacketHandshakeReject:
		return fmt.Errorf("connection refused, check client version")
	default:
		return fmt.Errorf("unknown packet type wth: %d", responsePacket.Type)
	}
}

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
		err := SendPacket(conn, protocol.Packet{
			Type: protocol.PacketHandshakeReject,
		})
		if err != nil {
			return fmt.Errorf("failed to send handshake reject")
		}
		return fmt.Errorf("invalid handshake version")
	}

	err = SendPacket(conn, protocol.Packet{
		Type: protocol.PacketHandshakeAck,
	})
	if err != nil {
		return fmt.Errorf("failed to write handshake ack")
	}
	return nil
}
