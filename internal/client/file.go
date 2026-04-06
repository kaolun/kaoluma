package client

import (
	"bytes"
	"encoding/binary"
	"io"
	"kaoluma/internal/core/protocol"
	"math/rand"
	"os"
)

const chunkSize = 50 * 1024

func SendFile(c *protocol.Client, path string, rng *rand.Rand) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	size := info.Size()
	name := info.Name()

	//TODO LIKE FRFR this is okay but it coullddddd break if rlly rlly unlucky so maybe hash or something later idfk
	fileID := rng.Uint32()

	startData, err := buildFileStart(fileID, size, name)
	if err != nil {
		return err
	}
	c.SendQueue <- protocol.Packet{
		Type: protocol.PacketFileStart,
		Data: startData,
	}

	buffer := make([]byte, chunkSize)
	chunkIndex := uint32(0)
	for {
		n, err := f.Read(buffer)
		if n > 0 {
			chunkData := buffer[:n]

			chunkData, err := buildFileChunk(fileID, chunkIndex, chunkData)
			if err != nil {
				return err
			}
			packet := protocol.Packet{
				Type: protocol.PacketFileChunk,
				Data: chunkData,
			}
			c.SendQueue <- packet
			chunkIndex++
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}
	endData, err := buildFileEnd(fileID)
	if err != nil {
		return err
	}
	endPacket := protocol.Packet{
		Type: protocol.PacketFileClose,
		Data: endData,
	}
	c.SendQueue <- endPacket
	return nil
}

// i dont really like funcs for not multi-use things but this is way too confusing without them
func buildFileStart(fileID uint32, size int64, filename string) ([]byte, error) {
	buf := new(bytes.Buffer)
	nameBytes := []byte(filename)

	if err := binary.Write(buf, binary.BigEndian, fileID); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, uint64(size)); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, uint16(len(nameBytes))); err != nil {
		return nil, err
	}
	if _, err := buf.Write(nameBytes); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// TODO add error checking later yk
func buildFileChunk(fileID uint32, index uint32, data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.BigEndian, fileID); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, index); err != nil {
		return nil, err
	}
	if _, err := buf.Write(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
func buildFileEnd(fileID uint32) ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, fileID); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
