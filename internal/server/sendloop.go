package server

import (
	"kaoluma/internal/core/transport"
	"log"
)

func SendLoop(c *ClientInfo) {
	for {
		packet := <-c.SendQueue
		err := transport.SendPacket(c.Conn, packet)
		if err != nil {
			log.Println("failed to send packet:", err)
			continue
		}
	}
}
