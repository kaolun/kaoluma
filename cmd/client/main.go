package main

import (
	"kaoluma/internal/client"
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
	"time"
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

	c := &protocol.Client{
		Conn: conn,
	}

	go client.ReadLoop(c)

	for i := 0; i < 50; i++ {
		packet := protocol.Packet{
			Type: protocol.PacketPing,
		}
		err = transport.SendPacket(c.Conn, packet)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
}
