package network

import (
	"encoding/binary"
	"fmt"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
)

func InitiateHandshake(c *protocol.Client) error {
	versionBuf := make([]byte, 2)
	binary.BigEndian.PutUint16(versionBuf, protocol.Version)

	packet := protocol.Packet{
		Type: protocol.PacketHandshake,
		Data: versionBuf,
	}

	err := transport.SendPacket(c.Conn, packet)
	if err != nil {
		return fmt.Errorf("failed to send handshake: %d", err)
	}

	responsePacket, err := protocol.DecodePacket(c)
	if err != nil {
		return fmt.Errorf("failed to decode handshake response: %d", err)
	}

	switch responsePacket.Type {
	case protocol.PacketHandshakeAck:
		return nil
	case protocol.PacketHandshakeReject:
		return fmt.Errorf("handshake rejected")
	default:
		return fmt.Errorf("unknown packet type wth: %d", responsePacket.Type)
	}
}

func PerformHandshake(c *protocol.Client) error {
	packet, err := protocol.DecodePacket(c)
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
		err := transport.SendPacket(c.Conn, protocol.Packet{
			Type: protocol.PacketHandshakeReject,
		})
		if err != nil {
			return fmt.Errorf("failed to send handshake reject")
		}
		return fmt.Errorf("invalid handshake version length")
	}

	err = transport.SendPacket(c.Conn, protocol.Packet{
		Type: protocol.PacketHandshakeAck,
	})
	if err != nil {
		return fmt.Errorf("failed to write handshake ack")
	}
	return nil
}
