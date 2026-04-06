package server

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"kaoluma/internal/core/protocol"
	"log"
	"os"
	"path/filepath"
)

func handleFileStart(s *Server, c *protocol.Client, packet protocol.Packet) error {
	fileID, fileName, fileSize, err := decodeFileStart(packet.Data)
	if err != nil {
		return err
	}
	tmpFile, err := os.CreateTemp("", "upload-*")
	if err != nil {
		return err
	}
	defer tmpFile.Close()
	s.IncomingFiles[fileID] = &protocol.IncomingFile{
		FileName:  fileName,
		FileSize:  fileSize,
		FileOwner: c.ID,
		Path:      tmpFile.Name(),
	}
	err = tmpFile.Close()
	if err != nil {
		return err
	}
	log.Println("filestart ts (ID:", fileID, ", Name:", fileName, ", Size:", fileSize, ")")
	return nil
}

func handleFileChunk(s *Server, c *protocol.Client, packet protocol.Packet) error {
	fileID, chunkIndex, chunkData, err := decodeFileChunk(packet.Data)
	if err != nil {
		return err
	}
	f, ok := s.IncomingFiles[fileID]
	if !ok {
		return fmt.Errorf("unknown id: %v", fileID)
	}
	tmpFile, err := os.OpenFile(f.Path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	offset := int64(chunkIndex) * protocol.ChunkSize
	if _, err := tmpFile.Seek(offset, 0); err != nil {
		return err
	}

	if _, err := tmpFile.Write(chunkData); err != nil {
		return err
	}

	log.Println("chunky (ID:", fileID, ", Index:", chunkIndex, ")")
	return nil
}

func handleFileClose(s *Server, c *protocol.Client, packet protocol.Packet) error {
	fileID, err := decodeFileEnd(packet.Data)
	if err != nil {
		return err
	}
	f, ok := s.IncomingFiles[fileID]
	if !ok {
		return fmt.Errorf("unknown id: %v", fileID)
	}
	if err := os.MkdirAll("./downloads", 0755); err != nil {
		return err
	}
	finalPath := "./downloads/" + f.FileName
	if err := os.Rename(f.Path, finalPath); err != nil {
		return err
	}
	log.Println("file done:", finalPath)
	delete(s.IncomingFiles, fileID)
	return nil
}

func decodeFileStart(data []byte) (uint32, string, uint64, error) {
	buf := bytes.NewReader(data)

	var fileID uint32
	var size uint64
	var nameLen uint16

	if err := binary.Read(buf, binary.BigEndian, &fileID); err != nil {
		return 0, "", 0, err
	}
	if err := binary.Read(buf, binary.BigEndian, &size); err != nil {
		return 0, "", 0, err
	}
	if err := binary.Read(buf, binary.BigEndian, &nameLen); err != nil {
		return 0, "", 0, err
	}

	nameBytes := make([]byte, nameLen)
	if _, err := buf.Read(nameBytes); err != nil {
		return 0, "", 0, err
	}
	name := string(nameBytes)

	return fileID, name, size, nil
}
func decodeFileChunk(data []byte) (uint32, uint32, []byte, error) {
	buf := bytes.NewReader(data)

	var fileID uint32
	var chunkIndex uint32
	var chunkData []byte

	if err := binary.Read(buf, binary.BigEndian, &fileID); err != nil {
		return 0, 0, nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &chunkIndex); err != nil {
		return 0, 0, nil, err
	}

	chunkData, err := io.ReadAll(buf)
	if err != nil {
		return 0, 0, nil, err
	}

	return fileID, chunkIndex, chunkData, nil
}
func decodeFileEnd(data []byte) (uint32, error) {
	buf := bytes.NewReader(data)

	var fileID uint32

	if err := binary.Read(buf, binary.BigEndian, &fileID); err != nil {
		return 0, err
	}

	return fileID, nil
}

// TODO remove this with ui or whatever the hell idek this part i dont get its too os'y
func getDownloadsDir() string {
	if dir, ok := os.LookupEnv("USERPROFILE"); ok {
		return filepath.Join(dir, "Downloads")
	}
	if dir, ok := os.LookupEnv("HOME"); ok {
		return filepath.Join(dir, "Downloads")
	}
	return "./downloads"
	//jst in case
}
