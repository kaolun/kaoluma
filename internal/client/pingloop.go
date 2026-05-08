package client

import (
	"kaoluma/internal/core/protocol"
	"time"
)

func PingLoop(c *Client) {
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
	return nil
}
