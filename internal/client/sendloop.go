package client

import (
	"kaoluma/internal/core/transport"
)

func SendLoop(c *Client) {
	for {
		select {
		case <-c.Done:
			return
		case packet, ok := <-c.SendQueue:
			if !ok {
				return
			}
			err := transport.SendPacket(c.Conn, packet)
			if err != nil {
				c.Log("Failed to send packet:" + err.Error())
				return
			}
		}
	}
}
