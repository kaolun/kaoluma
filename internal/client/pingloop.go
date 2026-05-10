package client

import (
	"kaoluma/internal/core/protocol"
	"time"
)

func PingLoop(c *Client) {
	ping := protocol.Packet{Type: protocol.PacketPing}
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-c.Done:
			return
		case <-ticker.C:
			if !c.UI.Online {
				return
			}
			c.Mu.RLock()
			since := time.Since(c.LastPong)
			c.Mu.RUnlock()
			if since > 10*time.Second {
				c.UI.Mu.Lock()
				c.UI.Online = false
				c.UI.Mu.Unlock()
				c.Log("Connection lost")
				return
			}
			c.SendQueue <- ping
		}
	}
}

func handlePong(c *Client, packet protocol.Packet) error {
	c.Mu.Lock()
	c.LastPong = time.Now()
	c.Mu.Unlock()
	return nil
}
