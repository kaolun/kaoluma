package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
)

func main() {
	port := ":8080"
	conn, err := client.Connect(port)
	if err != nil {
		log.Fatal("connection failed:", err)
	}
	err = transport.InitiateHandshake(conn)
	if err != nil {
		log.Fatal("Handshake failed:", err)
	}

	randomBS := []byte{1, 2, 3, 4, 5, 6, 7, 6, 7, 6, 7, 6, 7}
	packet := protocol.Packet{
		Type: protocol.PacketEcho,
		Data: randomBS,
	}

	err = transport.SendPacket(conn, packet)
	if err != nil {
		log.Fatal("aw shucks:", err)
	}

	incpacket, err := protocol.DecodePacket(conn)
	log.Printf("type: %d \n data: %d", incpacket.Type, incpacket.Data)
}
