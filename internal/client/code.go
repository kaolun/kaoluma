package client

import "kaoluma/internal/core/protocol"

func (c *Client) RequestJoinCode() {
	request := protocol.Packet{
		TargetID: 0,
		Type:     protocol.JoinCodeRequest,
	}
	c.SendQueue <- request
	c.Log("code requested")
}

func (c *Client) HandleJoinCode(p protocol.Packet) error {
	code := string(p.Data)
	c.Mu.Lock()
	c.JoinCode = code
	c.Mu.Unlock()
	c.Log("New joincode:" + code)
	return nil
}
