package handlers

import (
	"bytes"
	"encoding/binary"
	"kaoluma/internal/core/protocol"
	"log"
)

func HandleFileStart(c *protocol.Client, packet protocol.Packet) error {

	buf := bytes.NewReader(packet.Data)

	var fileID uint32
	var size uint64
	var nameLen uint16

	binary.Read(buf, binary.BigEndian, &fileID)
	binary.Read(buf, binary.BigEndian, &size)
	binary.Read(buf, binary.BigEndian, &nameLen)
	nameBytes := make([]byte, nameLen)
	buf.Read(nameBytes)

	name := string(nameBytes)

	log.Println("filestart ts ( ID:", fileID, ", Name:", name, ", Size:", size, ")")
	return nil
}
