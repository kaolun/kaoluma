package client

import (
	"kaoluma/internal/core/transport"
)

func SendLoop(c *Client) {
	for {
		select {
		case <-c.Done:
			return
		case packet := <-c.SendQueue:
			err := transport.SendPacket(c.Conn, packet)
			if err != nil {
				c.Log("Failed to send packet:" + err.Error())
				continue
			}
		}
	}
}
