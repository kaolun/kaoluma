package server

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
	"net"
	"time"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	client := &protocol.Client{
		Conn: conn,
	}

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	err := transport.PerformHandshake(client.Conn)
	if err != nil {
		log.Println("handshake failed:", err)
		return
	}

	for {
		err := conn.SetReadDeadline(time.Now().Add(30 * time.Second))
		if err != nil {
			log.Println(err)
			continue
		}

		packet, err := protocol.DecodePacket(client.Conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			log.Println("connection closed:", err)
			break
		}
		err = dispatch(client, packet)
		if err != nil {
			log.Println("error handling packet", err)
		}

	}
}
