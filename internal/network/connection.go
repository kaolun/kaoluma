package network

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/handlers"
	"log"
	"net"
	"time"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	err := PerformHandshake(conn)
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

		packet, err := protocol.DecodePacket(conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			log.Println("connection closed:", err)
			break
		}
		err = handlers.Dispatch(conn, packet)
		if err != nil {
			log.Println("error handling packet", err)
		}

	}
}
