package client

import (
	"kaoluma/internal/core/protocol"
	"net"
	"time"
)

func ReadLoop(c *Client) {
	for {
		select {
		case <-c.Done:
			return
		default:
		}
		c.Conn.SetReadDeadline(time.Now().Add(time.Second))

		packet, err := protocol.DecodePacket(c.Conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}
			select {
			case <-c.Done:
			default:
				close(c.Done)
				c.UI.Mu.Lock()
				c.UI.Online = false
				c.UI.Mu.Unlock()
				c.Log("Connection lost: " + err.Error())
			}
			return
		}

		err = dispatch(c, packet)
		if err != nil {
			c.Log("error handling packet:" + err.Error())
		}
	}
}
