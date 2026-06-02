package client

import "kaoluma/internal/core/protocol"

func (c *Client) RequestJoinCode() {
	c.SendQueue <- protocol.Packet{
		TargetID: 0,
		Type:     protocol.JoinCodeRequest,
	}
	c.Log("Code requested")
}

func (c *Client) HandleJoinCode(p protocol.Packet) error {
	code := string(p.Data)
	c.Mu.Lock()
	c.JoinCode = code
	c.Mu.Unlock()
	c.Log("New Joincode:" + code)
	return nil
}
