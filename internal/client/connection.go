package client

import (
	"fmt"
	"kaoluma/internal/core/transport"
	"net"
	"time"
)

func (c *Client) Connect() error {
	host := "localhost" + Port
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	c.Conn = conn
	c.Log("Connected")

	err = transport.InitiateHandshake(conn)
	if err != nil {
		return fmt.Errorf("handshake failed: %w", err)
	}
	c.Log("Handshake accepted")

	c.Mu.Lock()
	c.LastPong = time.Now()
	c.Done = make(chan struct{})
	c.Mu.Unlock()

	c.UI.Mu.Lock()
	c.UI.Online = true
	c.UI.Mu.Unlock()

	go ReadLoop(c)
	go SendLoop(c)
	go PingLoop(c)

	return err
}
