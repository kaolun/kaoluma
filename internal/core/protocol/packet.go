package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func DecodePacket(conn net.Conn) (Packet, error) {
	var typeBuf [1]byte
	_, err := io.ReadFull(conn, typeBuf[:])
	if err != nil {
		return Packet{}, err
	}

	var sIdBuff [4]byte
	_, err = io.ReadFull(conn, sIdBuff[:])
	if err != nil {
		return Packet{}, err
	}
	senderID := binary.BigEndian.Uint32(sIdBuff[:])
	var tIdBuff [4]byte
	_, err = io.ReadFull(conn, tIdBuff[:])
	if err != nil {
		return Packet{}, err
	}
	targetID := binary.BigEndian.Uint32(tIdBuff[:])

	var lengthBuf [4]byte
	_, err = io.ReadFull(conn, lengthBuf[:])
	if err != nil {
		return Packet{}, err
	}
	length := binary.BigEndian.Uint32(lengthBuf[:])
	if length > MaxPacketSize {
		return Packet{}, fmt.Errorf("packet exceeded size limit: %d bytes", length)
	}

	dataBuf := make([]byte, int(length))
	_, err = io.ReadFull(conn, dataBuf)
	if err != nil {
		return Packet{}, err
	}

	return Packet{
		Type:     PacketType(typeBuf[0]),
		SenderID: senderID,
		TargetID: targetID,
		Data:     dataBuf,
	}, nil
}

func EncodePacket(packet Packet) ([]byte, error) {
	if len(packet.Data) > MaxPacketSize {
		return nil, fmt.Errorf("packet too large: %d", len(packet.Data))
	}
	buf := make([]byte, 13+len(packet.Data))

	buf[0] = byte(packet.Type)
	binary.BigEndian.PutUint32(buf[1:5], uint32(packet.SenderID))
	binary.BigEndian.PutUint32(buf[5:9], uint32(packet.TargetID))
	binary.BigEndian.PutUint32(buf[9:13], uint32(len(packet.Data)))
	copy(buf[13:], packet.Data)

	return buf, nil
}
