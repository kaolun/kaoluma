package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

const MaxPacketSize = 64 * 1024

func DecodePacket(c *Client) (Packet, error) {
	var typeBuf [1]byte
	_, err := io.ReadFull(c.Conn, typeBuf[:])
	if err != nil {
		return Packet{}, err
	}

	var lengthBuf [4]byte
	_, err = io.ReadFull(c.Conn, lengthBuf[:])
	if err != nil {
		return Packet{}, err
	}
	length := binary.BigEndian.Uint32(lengthBuf[:])
	if length > MaxPacketSize {
		return Packet{}, fmt.Errorf("packet exceeded size limit: %d bytes", length)
	}

	dataBuf := make([]byte, int(length))
	_, err = io.ReadFull(c.Conn, dataBuf)
	if err != nil {
		return Packet{}, err
	}

	return Packet{
		Type: PacketType(typeBuf[0]),
		Data: dataBuf,
	}, nil
}

func EncodePacket(packet Packet) ([]byte, error) {
	if len(packet.Data) > MaxPacketSize {
		return nil, fmt.Errorf("packet too large: %d", len(packet.Data))
	}
	buf := make([]byte, 5+len(packet.Data))

	buf[0] = byte(packet.Type)
	binary.BigEndian.PutUint32(buf[1:5], uint32(len(packet.Data)))
	copy(buf[5:], packet.Data)

	return buf, nil
}
