package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const MaxPacketSize = 64 * 1024

func DecodePacket(conn net.Conn) (Packet, error) {
	var typeBuf [1]byte
	_, err := io.ReadFull(conn, typeBuf[:])
	if err != nil {
		return Packet{}, err
	}

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
		Type: PacketType(typeBuf[0]),
		Data: dataBuf,
	}, nil
}
