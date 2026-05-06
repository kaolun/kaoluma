package client

import (
	"kaoluma/internal/core/protocol"
	"log"
	"time"
)

func PingLoop(c *protocol.Client) {
	ping := protocol.Packet{Type: protocol.PacketPing}
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for range ticker.C {
		if len(c.SendQueue) == 0 {
			c.SendQueue <- ping
		}
	}
}

func handlePong(packet protocol.Packet) error {
	log.Println("ping success")
	return nil
}
