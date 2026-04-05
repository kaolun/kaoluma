package client

import (
	"kaoluma/internal/core/protocol"
	"kaoluma/internal/core/transport"
	"log"
)

func SendLoop(c *protocol.Client) {
	for {
		packet := <-c.SendQueue
		err := transport.SendPacket(c.Conn, packet)
		if err != nil {
			log.Println("failed to send packet:", err)
			continue
		}
	}
}
