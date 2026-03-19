package main

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/network"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	err = network.InitiateHandshake(conn)
	if err != nil {
		conn.Close()
		log.Fatal(err)
	}

	packet := protocol.Packet{
		Type: protocol.PacketEcho,
		Data: []byte{143, 67, 67, 67, 67, 143, 143, 192, 192, 12, 1, 1, 1, 1, 1, 1},
	}
	data, err := protocol.EncodePacket(packet)
	if err != nil {
		log.Println(err)
	}
	conn.Write(data)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	log.Println(buf[:n])

	conn.Close()
}
