package handlers

import (
	"bytes"
	"encoding/binary"
	"io"
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

	log.Println("filestart ts (ID:", fileID, ", Name:", name, ", Size:", size, ")")
	return nil
}

func HandleFileChunk(c *protocol.Client, packet protocol.Packet) error {
	buf := bytes.NewReader(packet.Data)

	var fileID uint32
	var chunkIndex uint32
	var chunkData []byte

	binary.Read(buf, binary.BigEndian, &fileID)
	binary.Read(buf, binary.BigEndian, &chunkIndex)

	chunkData, err := io.ReadAll(buf)
	if err != nil {
		return err
	}
	data := string(chunkData)

	log.Println("chunky (ID:", fileID, ", Index:", chunkIndex, ", Data:", data, ")")

	return nil
}

func HandleFileClose(c *protocol.Client, packet protocol.Packet) error {
	buf := bytes.NewReader(packet.Data)

	var fileID uint32

	binary.Read(buf, binary.BigEndian, &fileID)

	log.Println("file done")
	return nil
}
