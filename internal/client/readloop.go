package client

import (
	"kaoluma/internal/core/protocol"
	"log"
	"net"
	"time"
)

func ReadLoop(c *Client) {
	for {
		c.Conn.SetReadDeadline(time.Now().Add(time.Second))

		packet, err := protocol.DecodePacket(c.Conn)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			} else {
				log.Println("Connection closed:", err)
				break
			}
		}

		err = dispatch(c, packet)
		if err != nil {
			log.Println("error handling packet:", err)
		}
	}
}
